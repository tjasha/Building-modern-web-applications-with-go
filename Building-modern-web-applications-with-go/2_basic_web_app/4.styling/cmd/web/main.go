package main

import (
	"fmt"
	"net/http"

	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/handlers"
)

const portNumber = ":8080"

// we have to run it with "go run *.go" now

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) //we specify what to listen, in this case localhost on port 8080
}
