package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	databaseURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env not found")
	}

	return &Config{
		databaseURL: mustGetEnv("DATABASE_URL"),
	}

}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("env value is missing")
		log.Fatalf(" ENV '%s' need to be filled", key)
	}
	return value
}
