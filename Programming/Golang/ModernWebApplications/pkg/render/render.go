package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache:
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Get the requested template from cache
	t, isOk := tc[tmpl]
	if !isOk {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		fmt.Println(layouts)
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = templateSet
		fmt.Printf("Stored template: %s in cache\n", name)
	}

	return cache, nil
}
