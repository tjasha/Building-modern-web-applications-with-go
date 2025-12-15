package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

//for any function to respond to the request from the web browser, it needs 2 parameters - ResponseWriter and Request 
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl") //add name of the template that we want to render
} 

//second handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl") //add name of the template that we want to render
} 


func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) //we specify what to listen, in this case localhost on port 8080
}

// to use templates in go, we need to render them
func renderTemplates(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ) //build into the go system - takes path where templates are located and name of the template
	err := parsedTemplate.Execute(w, nil) //sending to where we want to execute it
	if err != nil {
		fmt.Println("Error pparsing templater: ", err)
		return
	}

}