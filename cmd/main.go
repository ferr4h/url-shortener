package main

import (
	"log"
	"net/http"
	configuration "url-shortener/config"
	"url-shortener/internal/url"
	"url-shortener/pkg/database"
	"url-shortener/pkg/middleware"
)

func main() {
	config := configuration.LoadConfig()

	cluster := database.NewCassandraCluster(config)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//Repositories
	urlRepository := url.NewUrlRepository(session)

	//Services
	urlService := url.NewUrlService(urlRepository)

	//Handlers
	router := http.NewServeMux()
	url.NewUrlHandler(router, config, urlService)

	server := &http.Server{
		Addr:    config.Host.Port,
		Handler: middleware.CORS(router),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
