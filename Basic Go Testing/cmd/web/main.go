package main

import (
	"basicgotesting/pkg/config"
	"basicgotesting/pkg/handlers"
	"basicgotesting/pkg/render"
	"fmt"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

func main() {
	//! Print the working directory here
	fmt.Println(os.Getwd())

	//! We want to load the template cache here
	tc, cacheErr := render.CreateTemplate()
	if cacheErr != nil {
		log.Fatal("Couldn't read the template cache when we wanted to")
	}

	//! Assign the template cache in the app config to be this
	app := config.AppConfig{TemplateCache: tc, UseCache: true}

	//! Set the global app config in render.go to have the same value
	render.NewTemplateCache(&app)

	//! Create a new repo
	Repo := handlers.NewRepo(&app)

	//! Create new handlers
	handlers.NewHandlers(Repo)

	//! Use the handlers and listen to the port
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	//! Print out progress to the data
	fmt.Printf("Starting the server on this port: %s", portNumber)

	//! Actually start the server
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("This is the error %s", err)
	}

}
