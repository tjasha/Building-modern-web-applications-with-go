package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Person struct {
	FirstName string `json:"first_name"` // if there is data from json format, this is exactly which field we want to associate with this field
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}



func main() {

	//case1 - read json

	myJson := `
	[
		{
			"first_name": "Clerk",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`


	var unmarshalled []Person   

	//Unmarshal receives 2 arguments. first is slice of bytes that is the content
	//second in interface that's a structure where you put content into 
	err := json.Unmarshal([]byte(myJson), &unmarshalled) 
	if err != nil { //we should check for errors
		log.Println("Error unmarshaling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)



	//case2 - write json from a struct

	var mySlice []Person

	var m1 Person 
	m1.FirstName = "Walley"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person 
	m2.FirstName = "Diana"
	m2.LastName = "Prine"
	m2.HairColor = "Black"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "      " ) //MarshalIndent is only formating json, in production you normally won't use it. prefix and indent are optional
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(newJson)
	fmt.Println(string(newJson)) //pars it to string

}
