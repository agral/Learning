package config

import "html/template"

// Holds the application-wide configuration.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
