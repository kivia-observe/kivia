package main

import (
	"log"

	"github.com/winnerx0/dyno/email_service/api"
	"github.com/winnerx0/dyno/email_service/internal/config"
	"github.com/winnerx0/dyno/email_service/internal/rabbitmq"
)

func main() {

	config := config.Load()
	
	client := rabbitmq.NewRabbitMQClient(config.RabbitMQURL)
	
	if err := client.SetupQueues(); err != nil {
		log.Fatal("Error setting up queues: ", err)
	}
	
	defer client.Close()

	server := api.NewServer(*config, client)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}

