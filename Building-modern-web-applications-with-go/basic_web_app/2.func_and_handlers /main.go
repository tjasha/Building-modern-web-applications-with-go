package main

import (
	"errors"
	"fmt"
	"net/http"
)

//i don't want this to be changed
const portNumber = ":8080"

//for any function to respond to the request from the web browser, it needs 2 parameters - ResponseWriter and Request 
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
} 

//second handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, err := fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
	if err != nil {
		fmt.Println(err)
	}
} 

//this is not a handler, it's just normal function
// it has lover case beginning so it will be only accessible inside the package
func addValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32)(float32, error){
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) //we specify what to listen, in this case localhost on port 8080
}