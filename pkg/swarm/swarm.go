package swarm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type swarm struct {
	addr string
}

type HealthStatusResponse struct {
	Status          string `json:"status"`
	Version         string `json:"version"`
	APIVersion      string `json:"apiVersion"`
	DebugAPIVersion string `json:"debugApiVersion"`
}

func (s *swarm) CheckNodeHealthAndReadiness() (*HealthStatusResponse, error) {
	url := "http://localhost:1635/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	m := new(HealthStatusResponse)

	_ = json.Unmarshal(body, &m)

	return m, nil
}

func (s *swarm) BuyPostageStamp() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetChequeBookBalance() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetChainState() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetNode() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetPeers() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetTransactions() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) Upload() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) Fetch() error {
	// TODO implement me
	panic("implement me")
}

func (s *swarm) GetBalance() error {
	// TODO implement me
	panic("implement me")
}

type Swarm interface {
	CheckNodeHealthAndReadiness() (*HealthStatusResponse, error)
	BuyPostageStamp() error
	GetChequeBookBalance() error
	GetChainState() error
	GetNode() error
	GetPeers() error
	GetTransactions() error
	Upload() error
	Fetch() error
	GetBalance() error
}

func New(addr string) Swarm {
	return &swarm{addr: addr}
}

// TODO: https://github.com/ethersphere/bee/tree/master/openapi
