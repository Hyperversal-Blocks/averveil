package main

import (
	"log"

	"github.com/hyperversal-blocks/averveil/av/cmd"
	"github.com/hyperversal-blocks/averveil/av/server"
)

func main() {
	desktopConfig, err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Init(desktopConfig); err != nil {
		log.Fatal(err)
	}
}
