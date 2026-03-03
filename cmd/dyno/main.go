package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/winnerx0/dyno/api"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/rabbitmq"
)

func main() {

	if err := godotenv.Load("../../.env"); err!= nil {
		log.Fatal("Error loading .env", err)
	}
	
	config := config.Load()
	
	rabbitMQClient := rabbitmq.NewRabbitMQClient(config.RabbitMQConnectionUrl)

	defer rabbitMQClient.Close()

	server := api.NewServer(*config, rabbitMQClient)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}
