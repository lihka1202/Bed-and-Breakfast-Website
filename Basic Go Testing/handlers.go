package main

import (
	"net/http"
)

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gohtml")
}
