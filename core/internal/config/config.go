package config

import (
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
		DBUrl:                 getEnv("DATABASE_URL", "postgresql://postgres:password@localhost:5000/kivia"),
		Port:                  getEnv("PORT", "8081"),
		JwtAccessTokenSecret:  getEnv("JWT_ACCESS_TOKEN_SECRET", "e057eb5d0604805f2809b58c892cc76ae7f6dffffee8e2d4f2807a4d8c86dfac"),
		JwtRefreshTokenSecret: getEnv("JWT_REFRESH_TOKEN_SECRET", "74b8fa85124b7cd11dfe34e33fdd6063711eab4d950af94affd0b50fe4ab3d85"),
		RabbitMQConnectionUrl: getEnv("RABBITMQ_CONNECTION_URL", "amqp://guest:guest@localhost:5672"),
		GoogleClientID:        getEnv("GOOGLE_CLIENT_ID", "890866835336-cef30b36umgfmehp7bgf5mnff8riri7m.apps.googleusercontent.com"),
		GoogleClientSecret:    getEnv("GOOGLE_CLIENT_SECRET", "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"),
		GoogleRedirectURL:     getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		GoogleFrontendURL:     getEnv("GOOGLE_FRONTEND_URL", "http://localhost:3000"),
	}

}

func getEnv(key string, defaultValue string) (value string) {

	if value = os.Getenv(key); value == "" {
		return defaultValue
	}

	return value

}
