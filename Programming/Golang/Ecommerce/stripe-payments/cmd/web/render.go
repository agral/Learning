package main

import (
	"embed"
	"html/template"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CssVersion      string
}

var functions = template.FuncMap{}

// go's embed templates:
var templateFS embed.FS
