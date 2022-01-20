package main

import (
	"flavioltonon/hmv/api"
	"log"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	defer server.Stop()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
