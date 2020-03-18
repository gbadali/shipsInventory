package mock

import (
	"time"

	"github.com/gbadali/shipsInventory/pkg/models"
)

var mockItem = &models.Item{
	ID:            1,
	ItemName:      "TestObject",
	Description:   "An Object for testing",
	NumOnHand:     12,
	LastInventory: time.Now(),
	Created:       time.Now(),
	Removed:       time.Time{},
	PartNum:       "C210-2019",
	Site:          "RL",
	Space:         "AMR",
	Drawer:        "C05",
}

type InventoryModel struct{}

func (m *InventoryModel) Insert(itemName, partNum, description, site, space, drawer string, numOnHand int) (int, error) {
	return 2, nil
}


func (m *InventoryModel) Get(id int) (*models.Item, error) {
	switch id {
	case 1:
		return mockItem, nil
	default:
		return nil, models.ErrNoRecord
	}
}


func (m *InventoryModel) Oldest() ([]*models.Item, error) {
	return []*models.Item{mockItem}, nil