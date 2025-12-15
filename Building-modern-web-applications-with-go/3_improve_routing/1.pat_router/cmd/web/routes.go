package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/config"
	"github.com/tjasha/Building-modern-web-aplications-with-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New() // we use external package, we installed it with "go get github.com/bmizerany/pat" from root

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}