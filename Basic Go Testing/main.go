package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Specify the port number

const portNum = ":8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gohtml")
}

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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start the server here and write to the console
	fmt.Println(fmt.Sprintf("App started on port: %s", portNum))

	_ = http.ListenAndServe(portNum, nil)

}
