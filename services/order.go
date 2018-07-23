package services

import (
	"../app"
	"../models"
)

// orderDAO specifies the interface of the order DAO needed by OrderService.
type orderDAO interface {
	// Get returns the order with the specified order ID.
	Get(rs app.RequestScope, id int) (*models.Order, error)
	// Count returns the number of orders.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of orders with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Order, error)
	// Create saves a new order in the storage.
	Create(rs app.RequestScope, order *models.Order) error
	// Update updates the order with given ID in the storage.
	Update(rs app.RequestScope, id int, order *models.Order) error
	// Delete removes the order with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// OrderService provides services related with orders.
type OrderService struct {
	dao orderDAO
}

// NewOrderService creates a new OrderService with the given order DAO.
func NewOrderService(dao orderDAO) *OrderService {
	return &OrderService{dao}
}

// Get returns the order with the specified the order ID.
func (s *OrderService) Get(rs app.RequestScope, id int) (*models.Order, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Order.
func (s *OrderService) Create(rs app.RequestScope, model *models.Order) (*models.Order, error) {
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Order_id)
}

// Update updates the order with the specified ID.
func (s *OrderService) Update(rs app.RequestScope, id int, model *models.Order) (*models.Order, error) {
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the order with the specified ID.
func (s *OrderService) Delete(rs app.RequestScope, id int) (*models.Order, error) {
	order, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return order, err
}

// Count returns the number of orders.
func (s *OrderService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the orders with the specified offset and limit.
func (s *OrderService) Query(rs app.RequestScope, offset, limit int) ([]models.Order, error) {
	return s.dao.Query(rs, offset, limit)
}