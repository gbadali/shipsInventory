package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create a middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// make a new mux
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/item/new", http.HandlerFunc(app.newItemForm))
	mux.Post("/item/new", http.HandlerFunc(app.newItem))
	mux.Get("/item/:id", http.HandlerFunc(app.showItem))

	// serve the staic files for css and js
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
