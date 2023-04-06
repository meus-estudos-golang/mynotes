package main

import (
	"net/http"

	"github.com/pauloa.junior/mynotes/internal/validator"
)

type userSignupFormData struct {
	Name     string
	Email    string
	Password string
	validator.Validator
}

func (app *application) userSignupForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupFormData{}
	app.render(w, http.StatusOK, "signup.tmpl.html", data)

	w.Write([]byte("Display user signup form"))
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := userSignupFormData{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
		return
	}

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
