package swarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type swarm struct {
	addr string
}

func (s *swarm) CheckNodeHealthAndReadiness() error {
	url := "http://localhost:1635/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(err)
	}
	m := make(map[string]string)
	_ = json.Unmarshal(body, &m)
	if m["status"] == "ok" {
		fmt.Println("is ok")
	}
	return nil
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

// TODO: https://github.com/ethersphere/bee/tree/master/openapi
