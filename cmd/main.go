package main

import (
	"log"
	"net/http"
	configuration "url-shortener/config"
	"url-shortener/internal/url"
	"url-shortener/pkg/database"
	"url-shortener/pkg/middleware"
)

func App(config *configuration.Config) http.Handler {
	cluster := database.NewCassandraCluster(config)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	redis := database.NewRedisClient(config)

	//Repositories
	urlRepository := url.NewUrlRepository(session, redis)

	//Services
	urlService := url.NewUrlService(urlRepository)

	//Handlers
	router := http.NewServeMux()
	url.NewUrlHandler(router, config, urlService)

	return middleware.CORS(router)
}

func main() {
	config := configuration.LoadConfig()
	app := App(config)

	server := &http.Server{
		Addr:    config.Host.Port,
		Handler: app,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
