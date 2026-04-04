package config

import (
	"fmt"
	"os"
)

type Config struct {
	BrevoApiKey string
	RabbitMQURL  string
	Port         string
	SenderEmail  string
	SenderName   string
}

func Load() *Config {

	return &Config{
		BrevoApiKey: getEnv("BREVO_API_KEY"),
		RabbitMQURL:  getEnv("RABBITMQ_CONNECTION_URL"),
		Port:         getEnv("PORT"),
		SenderEmail:  getEnv("SENDER_EMAIL"),
		SenderName:   getEnv("SENDER_NAME"),
	}

}

func getEnv(key string) (value string) {

	if value = os.Getenv(key); value == "" {
		panic(fmt.Sprintf("%s not in .env", key))
	}

	return value

}
