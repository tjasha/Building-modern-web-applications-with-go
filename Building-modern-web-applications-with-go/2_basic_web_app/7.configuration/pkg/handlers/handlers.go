package handlers

import (
	"net/http"

	"github.com/tjasha/modern_web_applications/pkg/config"
	"github.com/tjasha/modern_web_applications/pkg/render"
)

//Repository pattern: 

//repository used by handlers
var Repo *Repository

//repository type
type Repository struct {
	App *config.AppConfig
}

// create new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets repository for handlers
func NewHandlers(r *Repository){
	Repo = r
}


//we add receiver(m *Repository) to all handlers - this give them an access to the application variables
//Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl") //add name of the template that we want to render
} 

//About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl") //add name of the template that we want to render
} 
