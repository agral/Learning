package main

import (
	"GralLearning_ModernWebApps/pkg/handlers"
	"fmt"
	"net/http"
)

const PORT_NUMBER = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
