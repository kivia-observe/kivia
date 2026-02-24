package config

import (
	"os"
)

type Config struct {
	DBUrl string

	Port string
}

func Load() *Config {

	return &Config{
		DBUrl: getEnv("DATABASE_URL", "postgresql://postgres:2007@localhost:5432/dyno"),
		Port:  getEnv("PORT", "8000"),
	}

}

func getEnv(key string, defaultValue string) (value string) {

	if value = os.Getenv(key); value != "" {
		return value
	}

	return defaultValue

}
