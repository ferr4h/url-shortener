package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Host string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{DbConfig{Host: os.Getenv("Host")}}
}
