package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	configuration "url-shortener/config"
	"url-shortener/internal/url"
)

func TestAddUrlSuccess(t *testing.T) {
	config := configuration.LoadConfig()
	testServer := httptest.NewServer(App(config))
	defer testServer.Close()

	data, _ := json.Marshal(&url.CreateUrlRequest{
		Url: "https://google.com",
	})

	res, err := http.Post(testServer.URL+"/url", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 201 {
		t.Fatalf("Received non-201 response: %d\n", res.StatusCode)
	}
}

func TestAddUrlFail(t *testing.T) {
	config := configuration.LoadConfig()
	testServer := httptest.NewServer(App(config))
	defer testServer.Close()

	data, _ := json.Marshal(&url.CreateUrlRequest{
		Url: "not a link",
	})

	res, err := http.Post(testServer.URL+"/url", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 400 {
		t.Fatalf("Received non-400 response: %d\n", res.StatusCode)
	}
}
