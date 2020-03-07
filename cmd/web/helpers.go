package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// write an error and a stack trace to the error log
// then send a 500 code to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends the status code that it takes in as an int
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound just sends the 404 to the user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// render renders the template with the appropriate template data
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// Initialize a neww buffer.
	buf := new(bytes.Buffer)

	// try writing the template to the buffer to see if we get an error
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// if we didn't get an error we can write the buf to the http.ResponseWriter
	buf.WriteTo(w)
}
