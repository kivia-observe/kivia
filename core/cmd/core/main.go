package main

import (
	"log"
	"github.com/winnerx0/kivia/api"
	"github.com/winnerx0/kivia/internal/config"
	"github.com/winnerx0/kivia/internal/rabbitmq"

    "github.com/golang-migrate/migrate/v4"
    _"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbUrl string) {
    m, err := migrate.New(
        "file://migrations",
        dbUrl + "?sslmode=disable",
    )
    if err != nil {
        log.Fatal(err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }
}

func main() {

	config := config.Load()
	
	// RunMigrations(config.DBUrl)

	rabbitMQClient := rabbitmq.NewRabbitMQClient(config.RabbitMQConnectionUrl)

	defer rabbitMQClient.Close()

	server := api.NewServer(*config, rabbitMQClient)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}
