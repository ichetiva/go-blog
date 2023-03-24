package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresDatabase string

	SecretKey string
}

func NewConfig() *Config {
	// Load .env if exist
	godotenv.Load(".env")

	return &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresDatabase: os.Getenv("POSTGRES_DB"),

		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
