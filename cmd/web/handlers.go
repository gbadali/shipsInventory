package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gbadali/shipsInventory/pkg/models"
)

// GET /
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	i, err := app.inventory.Oldest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Items: i,
	})
}

// GET /item/:id
func (app *application) showItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	i, err := app.inventory.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Item: i,
	})
}

// GET /item/new
func (app *application) newItemForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

// POST /item/new
func (app *application) newItem(w http.ResponseWriter, r *http.Request) {
	// parse the form data
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the form data and put them in variales
	itemName := r.PostForm.Get("itemName")
	partNum := r.PostForm.Get("partNum")
	description := r.PostForm.Get("description")
	numOnHand := r.PostForm.Get("numOnHand")
	site := r.PostForm.Get("site")
	space := r.PostForm.Get("space")
	drawer := r.PostForm.Get("drawer")

	errors := make(map[string]string)

	if strings.TrimSpace(itemName) == "" {
		errors["itemName"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(itemName) > 100 {
		errors["itemName"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(description) == "" {
		errors["content"] = "This field cannot be blank"
	}

	num, err := strconv.Atoi(numOnHand)
	if err != nil {
		errors["numOnHand"] = "Cannot convert number on hand to int, setting it to 0"
		num = 0
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.inventory.Insert(itemName, partNum, description, site, space, drawer, num)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/item/%d", id), http.StatusSeeOther)
}
