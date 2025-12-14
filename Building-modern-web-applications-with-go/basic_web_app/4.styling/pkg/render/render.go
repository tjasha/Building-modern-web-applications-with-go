package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ) 
	err := parsedTemplate.Execute(w, nil) 
	if err != nil {
		fmt.Println("Error pparsing templater: ", err)
		return
	}

}