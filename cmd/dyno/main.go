package main

import (
	"log"

	"github.com/winnerx0/dyno/api"
	"github.com/winnerx0/dyno/internal/config"
)

func main() {

	config := config.Load()

	server := api.NewServer(*config)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}
