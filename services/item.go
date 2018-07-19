package services

import (
	"../app"
	"../models"
)

// itemDAO specifies the interface of the item DAO needed by ItemService.
type itemDAO interface {
	// Get returns the item with the specified item ID.
	Get(rs app.RequestScope, id int) (*models.Item, error)
	// Count returns the number of items.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of items with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Item, error)
	// Create saves a new item in the storage.
	Create(rs app.RequestScope, item *models.Item) error
	// Update updates the item with given ID in the storage.
	Update(rs app.RequestScope, id int, item *models.Item) error
	// Delete removes the item with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// ItemService provides services related with items.
type ItemService struct {
	dao itemDAO
}

// NewItemService creates a new ItemService with the given item DAO.
func NewItemService(dao itemDAO) *ItemService {
	return &ItemService{dao}
}

// Get returns the item with the specified the item ID.
func (s *ItemService) Get(rs app.RequestScope, id int) (*models.Item, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Item.
func (s *ItemService) Create(rs app.RequestScope, model *models.Item) (*models.Item, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.item_id)
}

// Update updates the item with the specified ID.
func (s *ItemService) Update(rs app.RequestScope, id int, model *models.Item) (*models.Item, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the item with the specified ID.
func (s *ItemService) Delete(rs app.RequestScope, id int) (*models.Item, error) {
	item, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return item, err
}

// Count returns the number of items.
func (s *ItemService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the items with the specified offset and limit.
func (s *ItemService) Query(rs app.RequestScope, offset, limit int) ([]models.Item, error) {
	return s.dao.Query(rs, offset, limit)
}