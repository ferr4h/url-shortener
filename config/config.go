package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host HostConfig
	Db   DbConfig
}

type HostConfig struct {
	Port string
}

type DbConfig struct {
	Host string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{HostConfig{Port: os.Getenv("PORT")}, DbConfig{Host: os.Getenv("CASSANDRA_HOST")}}
}
