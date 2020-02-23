package mysql

import (
	"database/sql"

	"github.com/gbadali/shipsInventory/pkg/models"
)

type InventoryModel struct {
	DB *sql.DB
}

func (m *InventoryModel) Insert(itemName, partNum, description, numOnHand string) (int, error) {
	return 0, nil
}

func (m *InventoryModel) Get(id int) (*models.Item, error) {
	return nil, nil
}

func (m *InventoryModel) Oldest() ([]*models.Snippet, error) {
	return nil, nil
}
