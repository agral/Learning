package main

import (
	"gochat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	const port = ":8080"
	log.Printf("Starting the server on port %s\n", port)
	_ = http.ListenAndServe(port, mux)
}
