package main

import (
	"fautest/handlers"
	"net/http"
)

func main() {
	router := handlers.New()
	server := http.Server{Addr: ":8080", Handler: router}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
