package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// Holds the application-wide configuration.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	IsProd        bool
	Session       *scs.SessionManager
}
