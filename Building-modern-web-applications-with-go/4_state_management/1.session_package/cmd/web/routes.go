package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/config"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New() // we use external package, we installed it with "go get github.com/bmizerany/pat" from root

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	//we can easily add some middleware that is available in chi
	//we are adding middleware that is recovering application if there was fatal error
	mux.Use(middleware.Recoverer) 
	// we are calling middleware here
	mux.Use(NoSerf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}