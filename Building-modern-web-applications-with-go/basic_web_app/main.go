package main

import (
	"fmt"
	"net/http"
)

func main() {

	// this handler is connecting to the web, checking path name - what user is typing in the web, homepage (/)
	// function that include writer = handler for response and request = pointer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ 
		n, err := fmt.Fprintf(w, "Hello world!") //we want to print this in the browser not in the console
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number ob bytes written: %d", n)) //using Sprintf to print int
	})

	_ = http.ListenAndServe(":8080", nil) //we specify what to listen, in this case localhost on port 8080
}