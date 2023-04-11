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

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.authenticate)

	// Unprotected rotes
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/note/view/:id", dynamic.ThenFunc(app.noteView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignupForm))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLoginForm))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLogin))

	protected := dynamic.Append(app.requireAuthentication)

	// Protected routes
	router.Handler(http.MethodGet, "/note/create", protected.ThenFunc(app.noteCreateForm))
	router.Handler(http.MethodPost, "/note/create", protected.ThenFunc(app.noteCreate))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogout))

	standard := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	return standard.Then(router)
}
