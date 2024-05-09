package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type AppConfig struct {
	UseCache      bool //! Needs to be in capital otherwise we cannot use UseCache outside this file
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
