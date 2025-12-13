package main

import (
	"log"

	helper "github.com/tjasha/packageTestProgram/Helper"
)

func main() {

	//library
	log.Println("Hello")

	var myVar helper.SomeType
	myVar.TypeName = "Jut name"
	log.Println(myVar.TypeName)


}

