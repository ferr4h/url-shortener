package request

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
)

func HandleBody[T any](r *http.Request) (*T, error) {
	body, err := Decode[T](r)
	if err != nil {
		return nil, err
	}
	err = Validate(body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func Decode[T any](r *http.Request) (T, error) {
	var payload T
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func Validate[T any](payload T) error {
	validate := validator.New()
	return validate.Struct(payload)
}
