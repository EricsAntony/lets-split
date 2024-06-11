package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// routes handles the routing of the application.
func (app *Application) routes() http.Handler {
	// Creates a middleware chain.
	middlewareChain := alice.New(app.Session.Enable)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Login))
	mux.Get("/home", http.HandlerFunc(app.Home))
	// mux.Get("/LoginUser", http.HandlerFunc(app.LoginUser))

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir)) // serve static files
	mux.Get("/static/", http.StripPrefix("/static", fileServer))  // strip static directory.

	return middlewareChain.Then(mux)
}
