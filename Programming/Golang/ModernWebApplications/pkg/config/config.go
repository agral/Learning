package config

import "html/template"

// Holds the application-wide configuration.
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
