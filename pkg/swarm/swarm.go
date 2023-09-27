package swarm

type swarm struct {
	addr string
}

func (s swarm) CheckNodeHealthAndReadiness() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) BuyPostageStamp() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) GetChequeBookBalance() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) GetChainState() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) GetNode() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) GetPeers() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) GetTransactions() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) Upload() error {
	// TODO implement me
	panic("implement me")
}

func (s swarm) Fetch() error {
	// TODO implement me
	panic("implement me")
}

type Swarm interface {
	CheckNodeHealthAndReadiness() error
	BuyPostageStamp() error
	GetChequeBookBalance() error
	GetChainState() error
	GetNode() error
	GetPeers() error
	GetTransactions() error
	Upload() error
	Fetch() error
}

func New() Swarm {
	return &swarm{}
}
