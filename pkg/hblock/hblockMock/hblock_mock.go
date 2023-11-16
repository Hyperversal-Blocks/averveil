package hblockMock

import (
	"context"
	"math/big"

	"github.com/hyperversal-blocks/averveil/pkg/hblock"
)

type hblockMock struct {
	getBalance func(ctx context.Context) (*big.Int, error)
}

func (g *hblockMock) GetBalance(ctx context.Context) (*big.Int, error) {
	return g.getBalance(ctx)
}

// Option is an option passed to New
type Option func(mock *hblockMock)

// New creates a new mock
func New(opts ...Option) hblock.Contract {
	bs := &hblockMock{}

	for _, o := range opts {
		o(bs)
	}

	return bs
}

func WithGetBalance(f func(ctx context.Context) (*big.Int, error)) Option {
	return func(mock *hblockMock) {
		mock.getBalance = f
	}
}
