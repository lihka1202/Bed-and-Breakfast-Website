package main

import (
	"fmt"
	"net/http"
)

// Specify the port number

const portNum = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start the server here and write to the console
	fmt.Println(fmt.Sprintf("App started on port: %s", portNum))

	_ = http.ListenAndServe(portNum, nil)

}
