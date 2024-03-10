package main

import (
	"fmt"
	"net/http"
)

// Specify the port number

const portNum = ":8080"

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start the server here and write to the console
	fmt.Println(fmt.Sprintf("App started on port: %s", portNum))

	_ = http.ListenAndServe(portNum, nil)

}
