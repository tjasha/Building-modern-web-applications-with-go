package main

import "log"


func main (){

	//only loop that exist in the go is for - while, do-while...

	//case1: 
	for i := 0; i <= 10; i++ {
		log.Println(i)
	} 

	//case2: 
	animals := []string{"dog", "cats", "fish", "horses"}

	for _, animal := range animals {
		log.Println(animal)
	}

	//case3: 
	houseAnimals := make(map[string]string)
	houseAnimals["dog"] = "Fido"
	houseAnimals["cat"] = "Fluffy"

	for animalType, animal := range houseAnimals {
		log.Println(animalType, animal)
	}


	//case4: 
	var firstLine = "Once upon a time"

	for i, l := range (firstLine) {
		log.Println(i, ":", l)
	}
	//if we are looping over the string, we get bytes, not letters! 


	//case5: 
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@email.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@email.com", 23})
	users = append(users, User{"Sally", "Brown", "sally@email.com", 38})
	users = append(users, User{"Alex", "SmiAndersonth", "alex@email.com", 34})


	for _, l := range (users) {
		log.Println(l.FirstName, l.LastName, l.Email)
	}

}