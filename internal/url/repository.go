package url

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/redis/go-redis/v9"
	"time"
)

const cacheExpiration = 10

var ctx = context.Background()

type UrlRepository struct {
	session *gocql.Session
	redis   *redis.Client
}

func NewUrlRepository(session *gocql.Session, redis *redis.Client) *UrlRepository {
	return &UrlRepository{session: session, redis: redis}
}

//TODO: error handling
//TODO: refactor repository

func (repo *UrlRepository) CreateUrl(hash, url string) error {
	createdAt := time.Now()
	err := repo.session.Query("INSERT INTO urls (hash, url, created_at) VALUES (?, ?, ?)", hash, url, createdAt).Exec()
	repo.setCacheEntry(hash, url)
	return err
}

func (repo *UrlRepository) GetUrl(hash string) (string, error) {
	var url string
	url, err := repo.getCacheByKey(hash)
	if err == redis.Nil {
		return url, nil
	}
	err = repo.session.Query("SELECT url FROM urls WHERE hash=?", hash).Scan(&url)
	return url, err
}

func (repo *UrlRepository) setCacheEntry(hash, url string) {
	err := repo.redis.Set(ctx, hash, url, time.Minute*cacheExpiration).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (repo *UrlRepository) getCacheByKey(hash string) (string, error) {
	result, err := repo.redis.Get(ctx, hash).Result()
	return result, err
}
