package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord means we couldn't find the item in the inventory
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials - incorrect email or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail - email is already registered in the database
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

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

// User struct stores the data pertainin to the user
// TODO: implment a admin column
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
