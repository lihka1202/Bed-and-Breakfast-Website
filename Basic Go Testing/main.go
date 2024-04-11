package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
}

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

func main() {
	//! Use the handlers and listen to the port
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//! Print out progress to the data
	fmt.Printf("Starting the server on this port: %s", portNumber)

	//! Actually start the server
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("This is the error %s", err)
	}
}
