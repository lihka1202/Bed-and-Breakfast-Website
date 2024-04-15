package handlers

import (
	"basicgotesting/pkg/config"
	"basicgotesting/pkg/models"
	"basicgotesting/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//! Get some data here
	stringMap := map[string]string{
		"test": "Hello World",
	}
	//! Pass as a pointer to preserve usage and save time
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{StringMap: stringMap})
}

// About is the about page handler
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{})
}
