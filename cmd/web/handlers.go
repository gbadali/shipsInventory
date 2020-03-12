package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	itemName := r.PostForm.Get("itemName")
	partNum := r.PostForm.Get("partNum")
	description := r.PostForm.Get("description")
	numOnHand := r.PostForm.Get("numOnHand")
	site := r.PostForm.Get("site")
	space := r.PostForm.Get("space")
	drawer := r.PostForm.Get("drawer")

	num, err := strconv.Atoi(numOnHand)
	if err != nil {
		app.clientError(w, http.StatusTeapot)
		return
	}

	id, err := app.inventory.Insert(itemName, partNum, description, site, space, drawer, num)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/item/:%d", id), http.StatusSeeOther)
}
