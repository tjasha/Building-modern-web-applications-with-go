package handlers

import (
	"net/http"

	"github.com/tjasha/moredn_web_applications/pkg/render"
)

// all handlers are in the same file

//Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl") //add name of the template that we want to render
} 

//About page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl") //add name of the template that we want to render
} 
