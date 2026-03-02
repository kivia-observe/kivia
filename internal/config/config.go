package config

import (
	"fmt"
	"os"
	
)

type Config struct {
	DBUrl string

	Port string

	JwtAccessTokenSecret string

	JwtRefreshTokenSecret string
	
	RabbitMQConnectionUrl string
}

func Load() *Config {

	return &Config{
		DBUrl:                 getEnv("DATABASE_URL"),
		Port:                  getEnv("PORT"),
		JwtAccessTokenSecret:  getEnv("JWT_ACCESS_TOKEN_SECRET"),
		JwtRefreshTokenSecret: getEnv("JWT_REFRESH_TOKEN_SECRET"),
		RabbitMQConnectionUrl: getEnv("RABBITMQ_CONNECTION_URL"),
	}

}

func getEnv(key string) (value string) {

	if value = os.Getenv(key); value == "" {
		panic(fmt.Sprintf("%s not in .env", key))
	}

	return value

}
