package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// to use templates in go, we need to render them
func renderTemplates(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ) //build into the go system - takes path where templates are located and name of the template
	err := parsedTemplate.Execute(w, nil) //sending to where we want to execute it
	if err != nil {
		fmt.Println("Error pparsing templater: ", err)
		return
	}

}