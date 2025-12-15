package main

import (
	"log"
	"time"
)

var s = "seven"

// capital letter is making variable visible outside of the package
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}


func main(){


	// case1: variable shadowing 

	var s2 = "six"

	// s := "eight" // don't do that! it can be easily confused

	log.Println("s1 is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")


	// types 

	user := User {
		FirstName: "John",
		LastName: "Doe",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)



}

func saySomething(s3 string) (string, string) {
	log.Println("s from the func", s)
	return s3, "something"
}
