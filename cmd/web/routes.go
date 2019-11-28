package main

import (
	"net/http"

	"github.com/markbates/pkger"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.defaultMiddleware(app.home))
	mux.HandleFunc("/resume", app.defaultMiddleware(app.resume))

	fileServer := http.FileServer(pkger.Dir("/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

func (app *application) defaultMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return app.recoverPanic(app.logRequest(secureHeaders(next)))
}
