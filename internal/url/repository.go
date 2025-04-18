package url

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type UrlRepository struct {
	session *gocql.Session
}

func NewUrlRepository(session *gocql.Session) *UrlRepository {
	return &UrlRepository{session: session}
}

func (repo *UrlRepository) CreateUrl(hash, url string) error {
	createdAt := time.Now()
	err := repo.session.Query("INSERT INTO urls (hash, url, created_at) VALUES (?, ?, ?)", hash, url, createdAt).Exec()
	return err
}

func (repo *UrlRepository) GetUrl(hash string) (string, error) {
	var url string
	err := repo.session.Query("SELECT url FROM urls WHERE hash=?", hash).Scan(&url)

	var results []struct {
		URL  string
		Hash string
	}
	iter := repo.session.Query("SELECT hash, url FROM urls").Iter()
	defer iter.Close()
	var URL string
	var h string
	for iter.Scan(&h, &URL) {
		results = append(results, struct {
			URL  string
			Hash string
		}{URL, h})
	}
	fmt.Println(results)

	return url, err
}
