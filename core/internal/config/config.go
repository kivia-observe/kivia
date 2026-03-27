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

	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	GoogleFrontendURL  string
}

func Load() *Config {

	return &Config{
		DBUrl:                 getEnv("DATABASE_URL"),
		Port:                  getEnv("PORT"),
		JwtAccessTokenSecret:  getEnv("JWT_ACCESS_TOKEN_SECRET"),
		JwtRefreshTokenSecret: getEnv("JWT_REFRESH_TOKEN_SECRET"),
		RabbitMQConnectionUrl: getEnv("RABBITMQ_CONNECTION_URL"),
		GoogleClientID:        getEnv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:    getEnv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURL:     getEnv("GOOGLE_REDIRECT_URL"),
		GoogleFrontendURL:     getEnv("GOOGLE_FRONTEND_URL"),
	}

}

func getEnv(key string) (value string) {

	if value = os.Getenv(key); value == "" {
		panic(fmt.Sprintf("%s not in .env", key))
	}

	return value

}
