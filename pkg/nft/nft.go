package nft

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/node"
	"github.com/hyperversal-blocks/averveil/pkg/store"
)

type nft struct {
	logger *logrus.Logger
	node   *node.Node
	db     store.Store
}

type NFT interface {
	Mint() (common.Hash, error)
}

func New() NFT {
	return &nft{}
}

func (n *nft) Mint() (common.Hash, error) {
	return [32]byte{}, nil
}
