package config

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string

	ServerAddress   string
	DbUrl           string
	JWTSecret       string
	RefreshTokenTTL time.Duration

	ProviderApiKey     string
	ProviderBaseApiURL *url.URL

	PrivateKey any
}

var Cfg *Config

func init() {
	Cfg = NewConfig()
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePrivateKey, err := os.ReadFile(fmt.Sprintf("%s/keys/private", wd))
	if err != nil {
		panic(err)
	}

	privateKey := parsePrivateKey(filePrivateKey)

	providerBaseUrl, err := url.Parse(os.Getenv("PROVIDER_BASE_API_URL"))
	if err != nil {
		panic(err)
	}

	return &Config{
		Env: getEnv("APP_ENV"),

		ServerAddress: getEnv("SERVER_ADDRESS"),
		DbUrl:         getEnv("DB_URL"),
		JWTSecret:     getEnv("SECRET_KEY"),

		ProviderApiKey:     getEnv("PROVIDER_API_KEY"),
		ProviderBaseApiURL: providerBaseUrl,

		PrivateKey: privateKey,
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("key %s has no value", key))
	}

	return value
}

func parsePrivateKey(privateKey []byte) any {
	block, _ := pem.Decode(privateKey)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		panic(err)
	}

	return key
}
