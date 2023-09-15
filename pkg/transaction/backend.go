package transaction

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Backend is the minimum of blockchain backend functions we need.
type Backend interface {
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	BlockNumber(ctx context.Context) (uint64, error)
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	ChainID(ctx context.Context) (*big.Int, error)

	Close()
}

// IsSynced will check if we are synced with the given blockchain backend. This
// is true if the current wall clock is after the block time of last block
// with the given maxDelay as the maximum duration we can be behind the block
// time.
func IsSynced(ctx context.Context, backend Backend, maxDelay time.Duration) (bool, time.Time, error) {
	number, err := backend.BlockNumber(ctx)
	if err != nil {
		return false, time.Time{}, err
	}
	header, err := backend.HeaderByNumber(ctx, big.NewInt(int64(number)))
	if errors.Is(err, ethereum.NotFound) {
		return false, time.Time{}, nil
	}
	if err != nil {
		return false, time.Time{}, err
	}

	blockTime := time.Unix(int64(header.Time), 0)

	return blockTime.After(time.Now().UTC().Add(-maxDelay)), blockTime, nil
}

// WaitSynced will wait until we are synced with the given blockchain backend,
// with the given maxDelay duration as the maximum time we can be behind the
// last block.
func WaitSynced(ctx context.Context, backend Backend, maxDelay time.Duration) error {
	for {
		synced, blockTime, err := IsSynced(ctx, backend, maxDelay)
		if err != nil {
			return err
		}

		if synced {
			return nil
		}

		fmt.Println("waiting on sync, blockTime: ", blockTime)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}
