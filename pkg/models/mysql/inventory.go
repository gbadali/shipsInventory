package mysql

import (
	"database/sql"
	"errors"
	"fmt"

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
	fmt.Println("trying to get item with id ", id)
	stmt := `SELECT id, itemName, description, created, lastInventory, numOnHand, partNum, site, space, drawer FROM inventory
		WHERE removed is NULL AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	i := &models.Item{}

	err := row.Scan(&i.ID, &i.ItemName, &i.Description, &i.Created, &i.LastInventory,
		&i.NumOnHand, &i.PartNum, &i.Site, &i.Space, &i.Drawer)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return i, nil
}

func (m *InventoryModel) Oldest() ([]*models.Item, error) {
	return nil, nil
}
