package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host  HostConfig
	Db    DbConfig
	Cache CacheConfig
}

type HostConfig struct {
	Port string
}

type DbConfig struct {
	Host     string
	Keyspace string
}

type CacheConfig struct {
	Host string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		HostConfig{Port: os.Getenv("PORT")},
		DbConfig{Host: os.Getenv("CASSANDRA_HOST"), Keyspace: os.Getenv("CLUSTER_KEYSPACE")},
		CacheConfig{os.Getenv("REDIS_HOST")},
	}
}
