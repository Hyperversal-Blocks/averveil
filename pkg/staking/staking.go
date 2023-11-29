package staking

type staking struct {
}

type Staking interface {
}

func New() Staking {
	return &staking{}
}
