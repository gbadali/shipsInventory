package mysql

import (
	"database/sql"

	"github.com/gbadali/shipsInventory/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Insert a new user into the db
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate verifies if a user exists with the provided email address
// and password
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
