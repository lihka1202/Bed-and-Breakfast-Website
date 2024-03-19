package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Package level variable
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	// Check if the template even exits in tc

	tmpl, inMap := tc[t]

	//! If present inMap would be set to true
	if !inMap {
		//? Create the template
		log.Println("Creating a template as it hasn't been cached before")

		err := createTemplate(t)
		if err != nil {
			fmt.Println("Error encountered here: ", err)
			return
		}
	} else {
		//! State that the template is present and being used
		log.Println("Using a cached template")
	}

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("There is an error while reading the template: ", err)
		return
	}
}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}

	//! Parse the template and then execute
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		fmt.Println("Template cannot be parsed due to", err)
	}
	//! Store it inside the holder
	tc[t] = tmpl

	//! No error then return as such
	return nil
}
