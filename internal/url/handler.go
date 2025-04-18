package url

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
	"url-shortener/config"
	"url-shortener/pkg/request"
	"url-shortener/pkg/response"
)

const hashLength = 6

type UrlHandler struct {
	config     *config.Config
	repository *UrlRepository
}

func NewUrlHandler(router *http.ServeMux, config *config.Config, repository *UrlRepository) {
	handler := UrlHandler{config: config, repository: repository}
	router.HandleFunc("POST /url", handler.CreateUrl())
	router.HandleFunc("GET /{alias}", handler.GetUrl())
}

func (handler UrlHandler) CreateUrl() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[CreateUrlRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hash := generateHash(req.Url, hashLength)
		err = handler.repository.CreateUrl(hash, req.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.WriteJSON(w, hash, http.StatusCreated)
	}
}

func (handler UrlHandler) GetUrl() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("alias")
		url, err := handler.repository.GetUrl(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
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
