package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type templateData struct {
	StringMap          map[string]string
	IntMap             map[string]int
	FloatMap           map[string]float32
	Data               map[string]interface{}
	CSRFToken          string
	Flash              string
	Warning            string
	Error              string
	IsAuthenticated    int
	API                string
	CssVersion         string
	StripeApiKey       string
	StripeApiSecretKey string
}

var functions = template.FuncMap{}

// NOTE: this is a special comment telling `embed` package where to look for templates.
//
//go:embed templates/*
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	// Read Stripe's API key from env variable.
	// This key is not provided directly, as it's a secret
	// that should not be exposed in the codebase.
	td.StripeApiKey = os.Getenv("STRIPE_API_PUBLISHABLE_KEY")
	if td.StripeApiKey == "" {
		app.errorLog.Println("Missing required environment variable: STRIPE_API_PUBLISHABLE_KEY")
		app.errorLog.Println("Stripe content will not be available")
	}
	// Read Stripe's API secret key from env variable.
	// This key is not provided directly, as it's a secret
	// that should not be exposed in the codebase.
	td.StripeApiSecretKey = os.Getenv("STRIPE_API_SECRET_KEY")
	if td.StripeApiSecretKey == "" {
		app.errorLog.Println("Missing required environment variable: STRIPE_API_SECRET_KEY")
		app.errorLog.Println("Stripe content will not be available")
	}
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request,
	page string, td *templateData, partials ...string) error {

	var t *template.Template
	var err error

	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)
	_, isFound := app.templateCache[templateToRender]

	if app.config.env == "production" && isFound {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(
	partials []string, page string, templateToRender string) (*template.Template, error) {

	var t *template.Template
	var err error

	// build the partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS,
			"templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS,
			"templates/base.layout.gohtml", templateToRender)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
