package main

import (
	"net/http"
)

// all handlers are in the same file

//Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl") //add name of the template that we want to render
} 

//About page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl") //add name of the template that we want to render
} 
