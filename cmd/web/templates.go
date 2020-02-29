package main

import (
	"github.com/gbadali/shipsInventory/pkg/models"
)

// TemplateData is here because we can only pass one struct into
// the template so we agregate it here.
type templateData struct {
	Item *models.Item
}
