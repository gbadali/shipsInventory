package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create a middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)

	// make a new mux
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/item/new", dynamicMiddleware.ThenFunc(app.newItemForm))
	mux.Post("/item/new", dynamicMiddleware.ThenFunc(app.newItem))
	mux.Get("/item/:id", dynamicMiddleware.ThenFunc(app.showItem))

	// serve the staic files for css and js
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
