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
	fmt.Println("Hello World")

	whatToSay string = "Hello Soon"

	fmt.Println(whatToSay)
}

func saySomething() string {
	return "nothing"
}