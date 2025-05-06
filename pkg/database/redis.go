package database

import (
	"github.com/redis/go-redis/v9"
	"url-shortener/config"
)

func NewRedisClient(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cache.Host,
		Password: "",
		DB:       0,
	})
	return rdb
}
