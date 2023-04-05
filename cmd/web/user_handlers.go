package main

import "net/http"

func (app *application) userSignupForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display user signup form"))
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new user"))
}

func (app *application) userLoginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display user login form"))
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate and login the user"))
}

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout the user"))
}
