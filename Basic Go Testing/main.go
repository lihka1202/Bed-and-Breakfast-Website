package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	n, err := fmt.Fprintf(w, "testing")

	if(err != nil) {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	//! Start the server here
	_ = http.ListenAndServe(":8080", nil)
}