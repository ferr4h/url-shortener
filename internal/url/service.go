package url

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

const hashLength = 6

type UrlService struct {
	repository *UrlRepository
}

func NewUrlService(repository *UrlRepository) *UrlService {
	return &UrlService{repository}
}

func (service UrlService) CreateUrl(url string) (string, error) {
	attempt := 0
	//Until unique hash is generated
	for {
		hash := generateHash(url, hashLength+attempt)
		existingUrl, err := service.repository.GetUrl(hash)
		if existingUrl == "" {
			err = service.repository.CreateUrl(hash, url)
			return hash, err
		}
		if existingUrl == url {
			return hash, nil
		}
		attempt++
	}
}

func (service UrlService) GetUrl(hash string) (string, error) {
	url, err := service.repository.GetUrl(hash)
	return url, err
}

func base62Encode(data []byte) string {
	str := base64.RawURLEncoding.EncodeToString(data)
	return strings.ReplaceAll(strings.ReplaceAll(str, "-", ""), "_", "")
}

func generateHash(url string, length int) string {
	hash := sha256.Sum256([]byte(url))
	b62 := base62Encode(hash[:])
	if len(b62) < length {
		return b62
	}
	return b62[:length]
}
