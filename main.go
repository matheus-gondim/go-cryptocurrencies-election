package main

import (
	"log"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/config"
	"github.com/matheus-gondim/go-cryptocurrencies-election/src/database"
	"github.com/matheus-gondim/go-cryptocurrencies-election/src/server"
)

func main() {
	configs := config.Init()

	db, err := database.NewConnection(&configs)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(db)

	server := server.NewCryptocurrencyElectionServer(db)

	if err := server.Run(configs.PORT); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
