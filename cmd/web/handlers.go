package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"github.com/pauloa.junior/mynotes/internal/models"
)

type noteCreateFormData struct {
	Title       string
	Content     string
	FieldErrors map[string]string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	notes, err := app.notes.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Notes = notes

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) noteView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	note, err := app.notes.GetById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Note = note

	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) noteCreateForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = noteCreateFormData{}

	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := noteCreateFormData{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 50 {
		form.FieldErrors["title"] = "This field cannot be more than 50 characters long"
	}

	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.notes.Insert(form.Title, form.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", id), http.StatusSeeOther)
}
