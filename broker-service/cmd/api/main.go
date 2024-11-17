package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	webPort = 80
)

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker server on port %s \n", strconv.Itoa(webPort))

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
