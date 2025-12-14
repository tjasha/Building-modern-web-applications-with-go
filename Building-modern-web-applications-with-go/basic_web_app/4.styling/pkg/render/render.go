package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, templ string) {
	//we need to add all templates here 
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ, "./templates/base.layout.tmpl") 
	err := parsedTemplate.Execute(w, nil) 
	if err != nil {
		fmt.Println("Error pparsing templater: ", err)
		return
	}

}