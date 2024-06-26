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
	Auth := alice.New(app.requireAuthenticatedUser)
	mux := pat.New()
	mux.Get("/login", http.HandlerFunc(app.Login))
	mux.Post("/login", http.HandlerFunc(app.LoginUser))
	mux.Get("/", Auth.ThenFunc(app.Home))
	mux.Get("/allusers", Auth.ThenFunc(app.AllUsers))
    mux.Get("/adduser", Auth.ThenFunc(app.AddUserform))
    mux.Post("/adduser", Auth.ThenFunc(app.AddUser))
	mux.Get("/logout", Auth.ThenFunc(app.Logout))

	fileServer := http.FileServer(http.Dir(app.Config.StaticDir)) // serve static files
	mux.Get("/static/", http.StripPrefix("/static", fileServer))  // strip static directory.

	return middlewareChain.Then(mux)
}
