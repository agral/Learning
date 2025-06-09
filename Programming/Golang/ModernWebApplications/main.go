package main

import (
	"fmt"
	"net/http"
)

const PORT_NUMBER = ":31337"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, World!")
}

func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the about page.")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
