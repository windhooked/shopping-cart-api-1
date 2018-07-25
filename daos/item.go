package daos

import (
	"../app"
	"../models"
	"strconv"
)

// ItemDAO persists data in database
type ItemDAO struct{}

func NewItemDAO() *ItemDAO {
	return &ItemDAO{}
}

// Get reads the item with the specified ID from the database.
func (dao *ItemDAO) Get(rs app.RequestScope, id int) (*models.Item, error) {
	// Check if item is cached in Redis otherwise use query to the DB
    cachedItem, err := models.FindItem(strconv.Itoa(id))
    if err == models.ErrNoItem {
		var item models.Item
		err := rs.Tx().Select().Model(id, &item)
		return &item, err
    } else if err != nil {
        return nil, err
    } else {
    	return cachedItem, err
    }
}

// Create saves a new item record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ItemDAO) Create(rs app.RequestScope, item *models.Item) error {
	item.Item_id = 0
	return rs.Tx().Model(item).Insert()
}

// Update saves the changes to an item in the database.
func (dao *ItemDAO) Update(rs app.RequestScope, id int, item *models.Item) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	item.Item_id = id
	return rs.Tx().Model(item).Exclude("Id").Update()
}

// Delete deletes an item with the specified ID from the database.
func (dao *ItemDAO) Delete(rs app.RequestScope, id int) error {
	item, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(item).Delete()
}

// Count returns the number of the item records in the database.
func (dao *ItemDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("item").Row(&count)
	return count, err
}

// Query retrieves the item records with the specified offset and limit from the database.
func (dao *ItemDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Item, error) {
	item := []models.Item{}
	err := rs.Tx().Select().OrderBy("item_id").Offset(int64(offset)).Limit(int64(limit)).All(&item)
	return item, err
}

