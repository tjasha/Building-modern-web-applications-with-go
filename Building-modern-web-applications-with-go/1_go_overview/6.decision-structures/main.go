package main

import "log"

func main(){


	//case1:
	isTrue := false

	// if isTrue {
	if isTrue == true {
		log.Println("isTrue is ", isTrue)
	} else{
		log.Println("isTrue is ", isTrue)
	}


	// case2
	cat := "not cat"

	// if isTrue {
	if "cat" == cat {
		log.Println("cat is ", cat)
	} else{
		log.Println("cat is ", cat)
	}


	//case3:

	myNum := 100
	isTrue2 := false

	if myNum > 99 && !isTrue2 {
		log.Println("condition is true")
	} else if myNum < 99 && !isTrue2 {
		log.Println("number is too small")
	} else {
		log.Println("condition is wrong")
	}

	// avoid too many else statements -> use switch
	

	//case4:
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("set to cat")
	case "dog":
		log.Println("set to dog")
	case "fish":
		log.Println("set to fish")
	default: 
		log.Println("something else")
	}

	//go breaks out as soon as it finds first matching case, it doesn't continue! 

}
