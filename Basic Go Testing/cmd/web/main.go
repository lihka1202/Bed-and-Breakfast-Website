package main

import (
	"basicgotesting/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

func main() {
	//! Print the working directory here
	fmt.Println(os.Getwd())

	//! Use the handlers and listen to the port
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//! Print out progress to the data
	fmt.Printf("Starting the server on this port: %s", portNumber)

	//! Actually start the server
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("This is the error %s", err)
	}

}
