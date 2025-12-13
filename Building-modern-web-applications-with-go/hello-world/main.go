package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello you")

	var whatToSay string
	var i int  

	whatToSay = "lalala"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to ", i)

	whatWasSaid, anotherThing := saySomething()
	fmt.Println("The function returned: ", whatWasSaid, anotherThing)
	
}

func saySomething() (string, string) {
	return "something", "somethingElse"
}