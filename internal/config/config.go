package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DBPort       string
	DBUser       string
	DBName       string
	DBPassword   string
	DBHost       string
	JWTSecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error("No .env file found")
	} else {
		slog.Info("Loaded .env file")
	}
	return &Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		DBPort:       os.Getenv("DATABASE_PORT"),
		DBUser:       os.Getenv("DATABASE_USER"),
		DBName:       os.Getenv("DATABASE_NAME"),
		DBPassword:   os.Getenv("DATABASE_PASSWORD"),
		DBHost:       os.Getenv("DATABASE_HOST"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
