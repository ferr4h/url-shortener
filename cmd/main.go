package main

import (
	"log"
	"net/http"
	configuration "url-shortener/config"
	"url-shortener/internal/url"
	"url-shortener/pkg/database"
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

	//Handlers
	router := http.NewServeMux()
	url.NewUrlHandler(router, config, urlRepository)

	server := &http.Server{
		Addr:    ":8081", //TODO: add to config
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
