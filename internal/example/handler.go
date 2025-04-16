package example

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/pkg/request"
	"url-shortener/pkg/response"
)

type ExampleHandler struct {
	config *config.Config
}

func NewExampleHandler(router *http.ServeMux, config *config.Config) {
	handler := &ExampleHandler{
		config: config,
	}
	router.HandleFunc("/", handler.Method())
}

func (h *ExampleHandler) Method() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[ExampleRequest](r)
		if err != nil {
			response.WriteJSON(w, err.Error(), 400)
			return
		}

		res := ExampleResponse{Line: req.Line}

		response.WriteJSON(w, res, 418)
	}
}
