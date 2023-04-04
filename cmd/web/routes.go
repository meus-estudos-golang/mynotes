package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.clientError(w, http.StatusNotFound)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/note/view/:id", app.noteView)
	router.HandlerFunc(http.MethodGet, "/note/create", app.noteCreateForm)
	router.HandlerFunc(http.MethodPost, "/note/create", app.noteCreate)

	standard := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	return standard.Then(router)
}
