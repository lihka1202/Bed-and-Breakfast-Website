package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/lihka1202/mainCode/pkg/config"
	"github.com/lihka1202/mainCode/pkg/handlers"
	"github.com/lihka1202/mainCode/pkg/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//! Print the working directory here
	fmt.Println(os.Getwd())

	//! We want to load the template cache here
	tc, cacheErr := render.CreateTemplate()
	if cacheErr != nil {
		log.Fatal("Couldn't read the template cache when we wanted to")
	}

	//! Assign the template cache in the app config to be this
	app = config.AppConfig{TemplateCache: tc, UseCache: true}

	//! Change this to true, when going in to production
	app.InProduction = false

	// ? Create the sessions package
	session = scs.New()

	//! Set the sessions time out to be 24 hours
	session.Lifetime = 24 * time.Hour

	//! cookie options
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	//! Make sure that sessions is held in appConfig
	//? No use as of now, but can be used in the future
	app.Session = session
	//! Set the global app config in render.go to have the same value
	render.NewTemplateCache(&app)

	//! Create a new repo
	Repo := handlers.NewRepo(&app)

	//! Create new handlers
	handlers.NewHandlers(Repo)

	//! Print out progress to the data
	fmt.Printf("Starting the server on this port: %s", portNumber)

	//! Actually start the server, pass in the handlers using PAT
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
