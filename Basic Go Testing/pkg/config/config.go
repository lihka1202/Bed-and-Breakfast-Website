package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool //! Needs to be in capital otherwise we cannot use UseCache outside this file
	TemplateCache map[string]*template.Template
}
