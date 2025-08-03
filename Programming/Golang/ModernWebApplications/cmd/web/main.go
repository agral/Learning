package main

import (
	"ModernWebApplications/pkg/config"
	"ModernWebApplications/pkg/handlers"
	"ModernWebApplications/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const PORT_NUMBER = ":31337"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Could not create the template cache")
		log.Fatal(err)
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
