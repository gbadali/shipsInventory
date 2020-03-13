package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gbadali/shipsInventory/pkg/forms"
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
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// POST /item/new
func (app *application) newItem(w http.ResponseWriter, r *http.Request) {
	// parse the form data
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("itemName", "description", "site", "space")
	form.MaxLength("itemName", 100)
	form.PermittedValues("site", "RL", "WH")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
	}

	num, err := strconv.Atoi(form.Get("numOnHand"))
	if err != nil {
		form.Errors.Add("numOnHand", "Could not convert Number on Hand to int")
		num = 0
	}

	id, err := app.inventory.Insert(
		form.Get("itemName"),
		form.Get("partNum"),
		form.Get("description"),
		form.Get("site"),
		form.Get("space"),
		form.Get("drawer"),
		num,
	)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/item/%d", id), http.StatusSeeOther)
}
