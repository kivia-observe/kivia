package main

import (
	"log"

	"github.com/winnerx0/dyno/api"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../../.env"); err!= nil {
		log.Fatal("Error loading .env", err)
	}

	config := config.Load()

	server := api.NewServer(*config)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}
