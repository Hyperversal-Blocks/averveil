package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/averveil/configuration"
	"github.com/hyperversalblocks/averveil/pkg/signer"
	"github.com/hyperversalblocks/averveil/pkg/transaction"
)

type Node struct {
	RpcClient *rpc.Client
	Signer    signer.Signer
	TxService transaction.Service
}

// InitNode initializes the node with passed configs
func InitNode(ctx context.Context, config configuration.Config, logger *logrus.Logger) (*Node, error) {
	rpcClient, err := Init(ctx, config.Chain.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("InitNode: unable to initialize rpcClient: %w", err)
	}

	owner, err := signer.New(config.Chain.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("InitNode: unable to initialize signer: %w", err)
	}

	backend := transaction.NewBackend(ethclient.NewClient(rpcClient))

	txService, err := transaction.NewTxService(logger, *backend, owner)
	if err != nil {
		return nil, fmt.Errorf("InitNode: unable to initialize transaction service: %w", err)
	}

	return &Node{
		RpcClient: rpcClient,
		Signer:    owner,
		TxService: txService,
	}, nil
}