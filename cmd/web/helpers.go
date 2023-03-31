package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	app.errorLog.Print(err.Error())
	http.Error(w, http.StatusText(status), status)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
