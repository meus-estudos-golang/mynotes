package main

import (
	"net/http"

	"github.com/pauloa.junior/mynotes/internal/validator"
)

type userSignupFormData struct {
	Name  string
	Email string
	validator.Validator
}

func (app *application) userSignupForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupFormData{}
	app.render(w, http.StatusOK, "signup.tmpl.html", data)

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
