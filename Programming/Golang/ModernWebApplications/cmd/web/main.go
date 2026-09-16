package main

import (
	"GralLearning_ModernWebApps/pkg/config"
	"GralLearning_ModernWebApps/pkg/handlers"
	"GralLearning_ModernWebApps/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const PORT_NUMBER = ":8080"

func main() {
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		// Instantiating all the templates has failed - just die at this point,
		// it makes no sense to proceed.
		log.Fatal("Cannot create the template cache.")
	}
	app.TemplateCache = templateCache

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
