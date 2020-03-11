package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gbadali/shipsInventory/pkg/models"
)

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

func (app *application) newItemForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}
func (app *application) newItem(w http.ResponseWriter, r *http.Request) {

	// Dummy data
	itemName := "O-Ring"
	description := "A purple O-Ring"
	numOnHand := 10
	partNum := "209-9088"
	site := "RL"
	space := "AMR"
	drawer := "C-06"

	id, err := app.inventory.Insert(itemName, partNum, description, site, space, drawer, numOnHand)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/item/:%d", id), http.StatusSeeOther)
}
