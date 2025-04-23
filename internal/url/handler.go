package url

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/pkg/request"
	"url-shortener/pkg/response"
)

type UrlHandler struct {
	config  *config.Config
	service *UrlService
}

func NewUrlHandler(router *http.ServeMux, config *config.Config, service *UrlService) {
	handler := UrlHandler{config: config, service: service}
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
		hash, err := handler.service.CreateUrl(req.Url)
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
		url, err := handler.service.GetUrl(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
}
