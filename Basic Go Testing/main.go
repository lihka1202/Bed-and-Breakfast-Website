package Basic_Go_Testing

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = "8080"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}q

func main() {
	//! Use the handlers and listen to the port
	http.HandleFunc("/", Home)
	http.HandleFunc("/", About)

	//! Print out progress to the data
	fmt.Printf("Starting the server on this port: %s", portNumber)

	//! Actually start the server
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalf("This is the error %s", err)
	}
}
