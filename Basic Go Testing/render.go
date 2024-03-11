package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate will allow us to render the template
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, fileErr := template.ParseFiles("./templates/" + tmpl)

	if fileErr != nil {
		fmt.Println("The templates are not being read properly, aborting")
		return
	}
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing the template: ", err)
		return
	}
}
