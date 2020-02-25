package mysql

import (
	"database/sql"

	"github.com/gbadali/shipsInventory/pkg/models"
)

type InventoryModel struct {
	DB *sql.DB
}

func (m *InventoryModel) Insert(itemName, partNum, description, site, space, drawer string, numOnHand int) (int, error) {
	stmt := `INSERT INTO inventory (itemName, description, created, lastInventory, numOnHand, partNum, site, space, drawer)
	VALUES(?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, itemName, description, numOnHand, partNum, site, space, drawer)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil

}

func (m *InventoryModel) Get(id int) (*models.Item, error) {
	return nil, nil
}

func (m *InventoryModel) Oldest() ([]*models.Item, error) {
	return nil, nil
}
