package main

import (
	"log"

	"averveil/cmd/server"
)

func main() {
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}
