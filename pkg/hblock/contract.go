package hblock

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/hyperversal-blocks/averveil/pkg/transaction"
)

type Contract interface {
	GetBalance(ctx context.Context) (*big.Int, error)
}

type contract struct {
	txService       transaction.Service
	contractAddress common.Address
	contractABI     abi.ABI
	owner           common.Address
	logger          *logrus.Logger
}

func (c *contract) GetBalance(ctx context.Context) (*big.Int, error) {
	callData, err := c.contractABI.Pack("balanceOf", c.owner)
	if err != nil {
		return nil, fmt.Errorf("unable to pack callData: %w", err)
	}

	result, err := c.txService.Call(ctx, &transaction.TxRequest{
		To:   &c.contractAddress,
		Data: callData,
	})
	if err != nil {
		return nil, fmt.Errorf("err calling tx service: %w", err)
	}

	results, err := c.contractABI.Unpack("balanceOf", result)
	if err != nil {
		return nil, fmt.Errorf("unable to unpack callData: %w", err)
	}

	if len(results) == 0 {
		return nil, errors.New("unexpected empty results")
	}

	return abi.ConvertType(results[0], new(big.Int)).(*big.Int), nil
}

func New(txService *transaction.Service, owner common.Address, logger *logrus.Logger, address common.Address, abi abi.ABI) Contract {
	return &contract{
		txService:       *txService,
		contractAddress: address,
		contractABI:     abi,
		owner:           owner,
		logger:          logger,
	}
}
