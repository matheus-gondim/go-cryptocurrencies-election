package main

import (
	"log"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/config"
	"github.com/matheus-gondim/go-cryptocurrencies-election/src/server"
)

func main() {
	configs := config.Init()

	server := server.NewCryptocurrencyElectionServer()

	if err := server.Run(configs.PORT); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
