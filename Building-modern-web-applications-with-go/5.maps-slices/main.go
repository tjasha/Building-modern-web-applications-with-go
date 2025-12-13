package main

import (
	"log"
	"sort"
)

type User struct{
	FirstName string
	LastName string
}
func main() {

	// case1:

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["anotherDog"] = "cassie"
	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])
	log.Println(myMap["anotherDog"])


	// case2:

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2
	log.Println(myMap2["First"])
	log.Println(myMap2["Second"])


	// case3: 

	myMap3 := make(map[string]User)

	me := User {
		FirstName: "Tjasa",
		LastName: "S",
	}

	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName)


	//maps will remain constant no matter if they are pass as a pointer or not. it will always change the value
	//they are NOT sorted! they will not be sorted int the way that they are put in or whatever. they are random
	// if you don't know what type you'll put in the map, you can use interface{}, but this is not recommended - you will also need to cast it back to what you need


	// case4 - slices: 

	var mySlice []string

	mySlice = append(mySlice, "Tjasa")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)


	//case5 

	var mySlice2 []int
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	log.Println(mySlice2)
	//slices are sorted, first value in will be on the first place


	//case6 - sorting

	sort.Ints(mySlice2)
	log.Println(mySlice2)


	//case7 - parts of slices 

	numbers := []int{1, 2, 3, 4, 5,6, 7, 8}
	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[5:])






}
