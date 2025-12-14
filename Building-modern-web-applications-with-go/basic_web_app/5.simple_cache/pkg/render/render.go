package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplates renders templates using html/template
func RenderTemplatesTest(w http.ResponseWriter, templ string) {
	//we need to add all templates here 
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ, "./templates/base.layout.tmpl") 
	err := parsedTemplate.Execute(w, nil) 
	if err != nil {
		fmt.Println("Error pparsing templater: ", err)
		return
	}
}

//to save templates to
var tc = make(map[string]*template.Template)

//this is simple cache enabling
func RenderTemplates (w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in our cache
	//we check in out map (tc) if template (t) exist in cache
	_, inMap := tc[t]
	if !inMap {
		//need to create new template
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		//we have the template in cache
		log.Println ("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

}

func createTemplateCache(t string) error {

	//we need to add all templates required to parse the web page !
	templates := []string{
		fmt.Sprintf("./templates/%s", t), 
		"./templates/base.layout.tmpl",   //winish with ,
	}

	//parse the templates and put them as individual strings
	tmpl, err := template.ParseFiles(templates...) 
	if err != nil {
		return err
	}

	//add template to cache (map)
	tc[t] = tmpl

	return nil
}