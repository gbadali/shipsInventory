package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Item struct {
	ID            int
	ItemName      string
	Description   string
	NumOnHand     int
	LastInventory time.Time
	Created       time.Time
	Removed       time.Time
	PartNum       string
	Site          string
	Space         string
	Drawer        string
}
