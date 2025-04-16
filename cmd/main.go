package main

import (
	"fmt"
	"log"
	"net/http"
	configuration "url-shortener/config"
	"url-shortener/internal/example"
	"url-shortener/pkg/database"
)

func main() {
	config := configuration.LoadConfig()

	cluster := database.NewCassandraCluster(config)
	fmt.Println(cluster)

	router := http.NewServeMux()
	example.NewExampleHandler(router, config)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
