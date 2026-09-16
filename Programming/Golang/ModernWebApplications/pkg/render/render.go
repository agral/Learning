package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Renders actual templates using html/template functionalities.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		// Instantiating all the templates has failed - just die at this point,
		// it makes no sense to proceed.
		log.Println("Failed to instantiate the template cache")
		log.Fatal(err)
	}

	// Try getting the requested template from the template cache:
	t, isOk := templateCache[tmpl]
	if !isOk {
		// Can't find the expected template in the cache - just die at this point.
		log.Printf("Failed to access the template: %s\n", tmpl)
		log.Fatal()
	}

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	// Finally render the template:
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	// Get all files matching: ./templates/*.page.tmpl:
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return myCache, err
	}

	// Process the found template files:
	for _, page := range pages {
		basename := filepath.Base(page)
		templateSet, err := template.New(basename).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[basename] = templateSet
	}

	log.Printf("Returning a cache: %v", myCache)
	return myCache, nil
}

var templateCache = make(map[string]*template.Template)
