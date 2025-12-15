package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/config"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/handlers"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/render"
)

const portNumber = ":8080"

// we have to run it with "go run *.go" now

func main() {

	var app config.AppConfig

	//i want to create template cache here
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false


	//this give render access to appConfig
	render.NewTemplates(&app)
	


	//create repository variable
	repo := handlers.NewRepo(&app)
	//create handlers and return variable back to handlers
	handlers.NewHandlers(repo)


	// we added this into routes
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil) //we specify what to listen, in this case localhost on port 8080

	//we add something that actually serves 
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	//we need to start a server
	err = srv.ListenAndServe()
	log.Fatal(err)
}
