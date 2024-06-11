package main

import (
	"log"
	"net/http"
)

// Home page for the application.
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
	}
	app.render(w, files, &templateData{
		UserId: app.Session.GetInt(r, "userId"),
	})
}

// Login shows the logins page.
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/login.page.tmpl",
		"ui/html/base.layout.tmpl",
	}
	app.render(w, files, nil)
}

// Login the user after authentication.
func (app *Application) LoginUser(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("username")
	password := r.FormValue("password")

	id, errAuth := app.User.Autenticate(user, password)
	if errAuth != nil {
		app.ErrorLog.Println(errAuth)
		log.Println(errAuth)
		return
	}
	app.Session.Put(r, "userId", id)
	log.Println("Done")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// AddUserForm renders the add user form.
func (app *Application) AddUserform(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/adduser.page.tmpl",
		"ui/html/base.layout.tmpl",
	}
	app.render(w, files, nil)
}

// AddUser adds a new user to the database.
func (app *Application) AddUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := app.User.InsertUser(name, email, password)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		log.Println("AddUser(): ", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout functionality.
func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r,"userId")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// AllUsers get the list of all users.
func (app *Application) AllUsers(w http.ResponseWriter, r *http.Request) {
    userlist, err := app.User.ListUsers()
    if err != nil {
        app.ErrorLog.Println(err.Error())
        log.Println("AllUsers(): ", err)
        return
    }
 
    files := []string{
        "ui/html/allusers.page.tmpl",
        "ui/html/base.layout.tmpl",
    }
    app.render(w, files, &templateData{
        UserList: userlist,
    })
}