package main

import (
	"log"

	"github.com/winnerx0/dyno/email_service/api"
)

func main() {

	server := api.NewServer()

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}

