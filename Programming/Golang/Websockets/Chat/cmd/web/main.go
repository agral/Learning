package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()
	const port = ":8080"
	log.Printf("Starting the server on port %s\n", port)
	_ = http.ListenAndServe(port, mux)
}
