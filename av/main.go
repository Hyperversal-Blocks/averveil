package main

import (
	"log"

	"github.com/hyperversal-blocks/averveil/av/cmd"
	"github.com/hyperversal-blocks/averveil/pkg/api"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Init(); err != nil {
		log.Fatal(err)
	}
}
