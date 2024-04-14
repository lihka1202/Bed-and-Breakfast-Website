package render

import (
	"html/template"
	"log"
	"net/http"
)

// ! Make a universal template cache to store parsed templates
var tc = make(map[string]*template.Template)

func RenderTemplateWorse(w http.ResponseWriter, tmpl string) {
	//! Note that using "./temaplates" is bad as it doesnt account for the working directory, in this case stick to this

	parsedTemplate, parseError := template.ParseFiles("./../../templates/"+tmpl, "./../../templates/base.layout.gohtml")
	if parseError != nil {
		log.Println("Something went wrong when we were trying to read the template")
		log.Printf("%v", parseError)
	}

	//! Execute this now
	execError := parsedTemplate.Execute(w, nil)

	if execError != nil {
		log.Println("There was an error while executing the template")
	}
}

// RenderTemplate checks if the template is in the map, if not adds it in
func RenderTemplate(w http.ResponseWriter, t string) {
	//! Check if template t is really in the map or not
	//? First argument is the value, second would be true if it exists or false otherwise
	_, inMap := tc[t]

	if !inMap {
		//! Create the map object in here
		err := createTemplate(t)
		if err != nil {
			log.Fatalf("Something is wrong hre")
		}

	} else {
		//! We know we using an existing template
		log.Println("Using cached template")
	}

	//! Execute it
	execErr := tc[t].Execute(w, nil)

	if execErr != nil {
		log.Fatal("Not working, there is an issue with execution")
	}
}

func createTemplate(t string) error {
	parsedTemplate, parseError := template.ParseFiles("./../../templates/"+t, "./../../templates/base.layout.gohtml")
	if parseError != nil {
		log.Println("Something went wrong when we were trying to read the template")
		return parseError
	}

	tc[t] = parsedTemplate

	//! return a  nil error
	return nil
}
