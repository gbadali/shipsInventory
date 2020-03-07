package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/gbadali/shipsInventory/pkg/models"
)

// TemplateData is here because we can only pass one struct into
// the template so we agregate it here.
type templateData struct {
	CurrentYear int
	Item        *models.Item
	Items       []*models.Item
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// here we parse the template files as well as add the functions we define
		// below so the templates know about the functions.  template.New creates a
		// new empty template set, then Funcs() registers the FuncMap and then we parse.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

// humanDate returns a nicely formatted string of the time
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// this is a string keyed map which acts as a lookup between the names
// of our custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}
