package transaction

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/signer"
)

var (
	// ErrTransactionReverted denotes that the transaction has been reverted.
	ErrTransactionReverted  = errors.New("transaction reverted")
	ErrUnknownTransaction   = errors.New("unknown transaction")
	ErrAlreadyImported      = errors.New("already imported")
	ErrTransactionCancelled = errors.New("transaction cancelled")
)

// TxRequest describes a request for a transaction that can be executed.
type TxRequest struct {
	To                   *common.Address // recipient of the transaction
	Data                 []byte          // transaction data
	GasPrice             *big.Int        // gas price or nil if suggested gas price should be used
	GasLimit             uint64          // gas limit or 0 if it should be estimated
	MinEstimatedGasLimit uint64          // minimum gas limit to use if the gas limit was estimated; it will not apply when this value is 0 or when GasLimit is not 0
	GasFeeCap            *big.Int        // adds a cap to maximum fee user is willing to pay
	Value                *big.Int        // amount of wei to send
	Description          string          // optional description
	GasTipBoost          int             // adds a tip for the miner for prioritizing transaction
	GasTipCap            *big.Int        // adds a cap to the tip
	Created              int64           // creation timestamp
	isCapped             bool
}

// Service is the service to send transactions. It takes care of gas price, gas
// limit and nonce management.
type Service interface {
	io.Closer
	// Send creates a transaction based on the request (with gasprice increased by provided percentage) and sends it.
	Send(ctx context.Context, request *TxRequest) (txHash common.Hash, err error)
	// Call simulate a transaction based on the request.
	Call(ctx context.Context, request *TxRequest) (result []byte, err error)
	// WaitForReceipt waits until either the transaction with the given hash has been mined or the context is cancelled.
	// This is only valid for transaction sent by this service.
	WaitForReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error)
	// CancelTransaction cancels a previously sent transaction by double-spending its nonce with zero-transfer one
	CancelTransaction(ctx context.Context, originalTxHash common.Hash) (common.Hash, error)
	// TransactionFee retrieves the transaction fee
	TransactionFee(ctx context.Context, txHash common.Hash) (*big.Int, error)
	// FilterLogs filters the events from contract
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) (*[]types.Log, error)
}

type TxService struct {
	wg     sync.WaitGroup
	lock   sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	logger    *logrus.Logger
	backend   WrappedBackend
	signer    signer.Signer
	sender    common.Address
	chainID   *big.Int
	rpcClient *rpc.Client
}

func (t *TxService) waitForPendingTx(txHash common.Hash) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		_, err := t.WaitForReceipt(t.ctx, txHash)
		if err != nil {
			if !errors.Is(err, ErrTransactionCancelled) {
				// TODO: add logger
				// e.g. t.logger.Error(err, "error while waiting for pending transaction", "tx", txHash)
				return
			} else {
				t.logger.Warning("pending transaction cancelled", "tx", txHash)
			}
		} else {
			t.logger.Debug("pending transaction confirmed", "tx", txHash)
		}
	}()
}

func (t *TxService) nextNonce(ctx context.Context) (uint64, error) {
	nonce, err := t.backend.PendingNonceAt(ctx, t.sender)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (t *TxService) nonceByBlock(ctx context.Context, sender common.Address, blockNum *big.Int) (uint64, error) {
	return t.backend.NonceAt(ctx, sender, blockNum)
}

// prepareTransaction creates a signable transaction based on a request.
func (t *TxService) prepareTransaction(ctx context.Context, request *TxRequest, nonce uint64) (tx *types.Transaction, err error) {

	gasLimit, err := t.backend.EstimateGas(ctx, ethereum.CallMsg{
		From: t.sender,
		To:   request.To,
		Data: request.Data,
	})
	if err != nil {
		return nil, err
	}

	gasLimit += gasLimit / 4 // add 25% on top
	if gasLimit < request.MinEstimatedGasLimit {
		gasLimit = request.MinEstimatedGasLimit
	}

	/*
		Transactions are EIP 1559 dynamic transactions where there are three fee related fields:
			1. base fee is the price that will be burned as part of the transaction.
			2. max fee is the max price we are willing to spend as gas price.
			3. max priority fee is max price want to give to the miner to prioritize the transaction.
		as an example:
		if base fee is 15, max fee is 20, and max priority is 3, gas price will be 15 + 3 = 18
		if base is 15, max fee is 20, and max priority fee is 10,
		gas price will be 15 + 10 = 25, but since 25 > 20, gas price is 20.
		notice that gas price does not exceed 20 as defined by max fee.
	*/

	if request.isCapped {
		gasFeeCap, gasTipCap, err := t.SuggestedFeeAndTip(ctx)
		if err != nil {
			return nil, err
		}

		if request.GasFeeCap == nil {
			request.GasFeeCap = gasFeeCap
		}

		if request.GasTipCap == nil {
			request.GasTipCap = gasTipCap
		}
	}
	return types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		ChainID:   t.chainID,
		To:        request.To,
		Value:     request.Value,
		Gas:       gasLimit,
		GasFeeCap: request.GasFeeCap,
		GasTipCap: request.GasTipCap,
		Data:      request.Data,
	}), nil
}

func (t *TxService) SuggestedFeeAndTip(ctx context.Context) (*big.Int, *big.Int, error) {
	var err error

	gasPrice, err := t.backend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}

	gasTipCap, err := t.backend.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}

	gasFeeCap := new(big.Int).Add(gasTipCap, gasPrice)

	t.logger.Debug("prepare transaction", "gas_price", gasPrice, "gas_max_fee", gasFeeCap, "gas_max_tip", gasTipCap)

	return gasFeeCap, gasTipCap, nil

}

func (t *TxService) FilterLogs(ctx context.Context, query ethereum.FilterQuery) (*[]types.Log, error) {
	filteredLogs, err := t.backend.FilterLogs(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to filter logs: %w", err)
	}

	return &filteredLogs, nil
}

func (t *TxService) Close() error {
	t.cancel()
	t.wg.Wait()
	return nil
}

func (t *TxService) Send(ctx context.Context, request *TxRequest) (txHash common.Hash, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	nonce, err := t.nextNonce(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := t.prepareTransaction(ctx, request, nonce)
	if err != nil {
		return common.Hash{}, err
	}

	signedTx, err := t.signer.SignTx(tx, t.chainID)
	if err != nil {
		return common.Hash{}, err
	}

	err = t.backend.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, err
	}

	txHash = signedTx.Hash()

	t.waitForPendingTx(txHash)

	return signedTx.Hash(), nil
}

func (t *TxService) Call(ctx context.Context, request *TxRequest) (result []byte, err error) {
	msg := ethereum.CallMsg{
		From:     t.sender,
		To:       request.To,
		Data:     request.Data,
		GasPrice: request.GasPrice,
		Gas:      request.GasLimit,
		Value:    request.Value,
	}
	data, err := t.backend.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *TxService) WaitForReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	return t.backend.TransactionReceipt(ctx, txHash)
}

func (t *TxService) CancelTransaction(ctx context.Context, originalTxHash common.Hash) (common.Hash, error) {
	receipt, err := t.WaitForReceipt(ctx, originalTxHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error getting the receipt from tx hash: %s with error: %w", originalTxHash, err)
	}

	gasFeeCap, gasTipCap, err := t.SuggestedFeeAndTip(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	gasTipCap = new(big.Int).Div(new(big.Int).Mul(big.NewInt(int64(10)+100), gasTipCap), big.NewInt(100))

	gasFeeCap.Add(gasFeeCap, gasTipCap)

	nonce, err := t.nonceByBlock(ctx, t.sender, receipt.BlockNumber)
	if err != nil {
		return [32]byte{}, fmt.Errorf("unable to get nonce by block: %w", err)
	}

	signedTx, err := t.signer.SignTx(types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		ChainID:   t.chainID,
		To:        &t.sender,
		Value:     big.NewInt(0),
		Gas:       21000,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      []byte{},
	}), t.chainID)
	if err != nil {
		return common.Hash{}, err
	}

	err = t.backend.SendTransaction(t.ctx, signedTx)
	if err != nil {
		return common.Hash{}, err
	}

	txHash := signedTx.Hash()

	t.waitForPendingTx(txHash)

	return txHash, err
}

func NewTxService(logger *logrus.Logger, backend WrappedBackend, signer signer.Signer) (Service, error) {
	ctx, cancel := context.WithCancel(context.Background())
	tx := &TxService{
		wg:      sync.WaitGroup{},
		lock:    sync.Mutex{},
		ctx:     ctx,
		cancel:  cancel,
		logger:  logger,
		backend: backend,
		signer:  signer,
	}
	return tx, nil
}

func (t *TxService) TransactionFee(ctx context.Context, txHash common.Hash) (*big.Int, error) {
	trx, _, err := t.backend.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return trx.Cost(), nil
}
