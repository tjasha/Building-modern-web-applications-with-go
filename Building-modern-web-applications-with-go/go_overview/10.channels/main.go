package main

import (
	"log"

	helpers "github.com/tjasha/packageTestProgram2/Helpers"
)


const numPool = 1000

func CalculateValue (intChan chan int){
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber //push value in the channel

}

func main() {

	intChan := make(chan int)  //this channel can only include int

	defer close(intChan) // execute this as soon as current function is done
						// we use it for everything that needs to be closed lik DB connections, channels and so on

	//we run go routine and write in the channel in it
	go CalculateValue(intChan)

	//we need to listen to response in the channel
	num := <-intChan //get value from the channel

	log.Println(num)


	//it's very useful to pass values from one package to another or one part of the program to another
}
