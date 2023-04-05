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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/note/view/:id", dynamic.ThenFunc(app.noteView))
	router.Handler(http.MethodGet, "/note/create", dynamic.ThenFunc(app.noteCreateForm))
	router.Handler(http.MethodPost, "/note/create", dynamic.ThenFunc(app.noteCreate))

	standard := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	return standard.Then(router)
}
