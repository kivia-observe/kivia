package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/winnerx0/dyno/api"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/rabbitmq"

    _"github.com/golang-migrate/migrate/v4"
    _"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// func RunMigrations(dbUrl string) {
//     m, err := migrate.New(
//         "file://migrations",
//         dbUrl + "?sslmode=disable",
//     )
//     if err != nil {
//         log.Fatal(err)
//     }

//     if err := m.Up(); err != nil && err != migrate.ErrNoChange {
//         log.Fatal(err)
//     }
// }

func main() {

	
	env := os.Getenv("APP_ENV")
	
	if env == "" {
		env = "local"
	}
	
	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(fmt.Sprintf("../../%s", envFile)); err != nil {
		env = "dev"
		log.Println("Dyno environment in ", env)
	}

	config := config.Load()
	
	// RunMigrations(config.DBUrl)

	rabbitMQClient := rabbitmq.NewRabbitMQClient(config.RabbitMQConnectionUrl)

	defer rabbitMQClient.Close()

	server := api.NewServer(*config, rabbitMQClient)

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
}
