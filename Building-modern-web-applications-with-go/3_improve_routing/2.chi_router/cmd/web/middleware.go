package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// we want to create middleware that is writing some text in console every time that sameone hits the page

//very often name of the parameter in middleware is called next
func WriteToConsole(next http.Handler) http.Handler {
	// handler with anonymous function 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r) //we have to say that we want to go further, to the next thing
	})
}

//creating noSerf token
func NoSerf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	//have to creat a cookie with som values, valid to all sites
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}