package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Package level variable
var tc = make(map[string]*template.Template)

// RenderTemplate will allow us to render the template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, fileErr := template.ParseFiles("./../../templates/"+tmpl, "./../../templates/base.layout.gohtml")

	if fileErr != nil {
		fmt.Println("The templates are not being read properly, aborting")
		fmt.Println(fileErr)
		return
	}
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing the template: ", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter, t string) {

	// Check if the template even exits in tc

	tmpl, inMap := tc[t]

	//! If present inMap would be set to true
	if !inMap {
		//? Create the template
	} else {
		//! State that the template is present and being used
		log.Println("Using a cached template")
	}

	err := tmpl.Execute(w, nil)
}
