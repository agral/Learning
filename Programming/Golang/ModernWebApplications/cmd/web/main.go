package main

import (
	"ModernWebApplications/pkg/config"
	"ModernWebApplications/pkg/handlers"
	"ModernWebApplications/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":31337"

func main() {
	var app config.AppConfig

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // good for debugging. but in production, this *has* to be set to true.

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Could not create the template cache")
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
