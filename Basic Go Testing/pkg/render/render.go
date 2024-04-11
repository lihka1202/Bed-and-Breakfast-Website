package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//! Note that using "./temaplates" is bad as it doesnt account for the working directory, in this case stick to this
	parsedTemplate, parseError := template.ParseFiles("./../../templates/" + tmpl)
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
