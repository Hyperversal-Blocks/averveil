package main

import (
	"log"

	"github.com/hyperversalblocks/averveil/cmd/server"
)

func main() {
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}
