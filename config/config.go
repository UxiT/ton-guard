package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DB_URL        string
	JWTSecret     string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS"),
		DB_URL:        getEnv("DB_URL"),
		JWTSecret:     getEnv("SECRET_KEY"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("key %s has no value", key))
	}

	return value
}
