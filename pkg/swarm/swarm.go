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

type Swarm interface {
}

func New() Swarm {
	return &swarm{}
}

func CheckNodeHealth() error {
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
