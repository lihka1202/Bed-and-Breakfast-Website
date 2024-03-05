package main

import (
	"fmt"
	"time"
)

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main () {
	user := User {
		FirstName: "Lol",
		LastName: "Finder",
		PhoneNumber: "9090",
		Age: 40,
	}

	fmt.Println(user.FirstName)
}