package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, parseError := template.ParseFiles("./templates/" + tmpl)
	if parseError != nil {
		log.Println("Something went wrong when we were trying to read the template")
	}

	//! Execute this now
	execError := parsedTemplate.Execute(w, nil)

	if execError != nil {
		log.Println("There was an error while executing the template")
	}
}
