package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create a middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	// make a new mux
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// Routes that have to do with items
	mux.Get("/item/new", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.newItemForm))
	mux.Post("/item/new", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.newItem))
	mux.Get("/item/:id", dynamicMiddleware.ThenFunc(app.showItem))
	// User Routes
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// serve the staic files for css and js
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
