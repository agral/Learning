package main

import (
	"log"
	"os"

	"github.com/agral/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// initialize Celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	cel.IsDebug = true

	app := &application{
		App: cel,
	}

	return app
}
