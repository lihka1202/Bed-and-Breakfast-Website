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

	//! Store the remote IP here, IPv4 or IPv6
	remoteIP := r.RemoteAddr

	// store the remote IP in the session using "remote_ip" as the key
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	//! Pass as a pointer to preserve usage and save time
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{StringMap: stringMap})
}

// About is the about page handler
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	//! Get the value from remoteIP
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")

	//! Place this session in the string map
	stringMap := map[string]string{}

	//! Place the values in the stringMap
	stringMap["test"] = "Hello World"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
