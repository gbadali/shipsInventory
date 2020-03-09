package main

import "net/http"

func (app *application) routes() http.Handler {
	// make a new mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/item", app.showItem)
	mux.HandleFunc("/item/add", app.addItem)

	// serve the staic files for css and js
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
