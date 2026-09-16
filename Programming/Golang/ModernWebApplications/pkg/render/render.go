package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// Renders actual templates using html/template functionalities.
func RenderTemplateNoCaching(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplateSimpleCaching(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check whether template is already in cache?
	_, isOk := templateCache[t]
	if !isOk {
		// Need to intantiate the template & store it in the map
		fmt.Println("Creating a new template")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Template already available in the map
		fmt.Println("Using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the template:
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add the parsed template to template cache:
	templateCache[t] = tmpl

	return nil
}

// Calls to RenderTemplate should from now on invoke RenderTemplateSimplified, not RenderTemplatePrevious.
func RenderTemplate(w http.ResponseWriter, t string) {
	RenderTemplateSimpleCaching(w, t)
}
