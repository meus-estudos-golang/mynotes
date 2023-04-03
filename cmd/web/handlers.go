package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/pauloa.junior/mynotes/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.clientError(w, http.StatusNotFound)
		return
	}

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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Outras"
	content := "Pagar contas"

	id, err := app.notes.Insert(title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("Nota criada com o ID = %d", id)
	w.Write([]byte("Create a new note"))
}
