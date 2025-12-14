package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// we have to run it with "go run *.go" now

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) //we specify what to listen, in this case localhost on port 8080
}
