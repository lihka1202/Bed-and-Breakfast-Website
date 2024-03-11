package main

import (
	"basicgotesting/pkg/handlers"
	"fmt"
	"net/http"
)

// Specify the port number

const portNum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Start the server here and write to the console
	fmt.Println(fmt.Sprintf("App started on port: %s", portNum))

	_ = http.ListenAndServe(portNum, nil)

}
