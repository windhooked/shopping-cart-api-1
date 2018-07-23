package services

import (
	"../app"
	"../models"
)

// customerDAO specifies the interface of the customer DAO needed by CustomerService.
type customerDAO interface {
	// Get returns the customer with the specified customer ID.
	Get(rs app.RequestScope, id int) (*models.Customer, error)
	// Count returns the number of customers.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of customers with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error)
	// Create saves a new customer in the storage.
	Create(rs app.RequestScope, customer *models.Customer) error
	// Update updates the customer with given ID in the storage.
	Update(rs app.RequestScope, id int, customer *models.Customer) error
	// Delete removes the customer with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// CustomerService provides services related with customers.
type CustomerService struct {
	dao customerDAO
}

// NewCustomerService creates a new CustomerService with the given customer DAO.
func NewCustomerService(dao customerDAO) *CustomerService {
	return &CustomerService{dao}
}

// Get returns the customer with the specified the cust ID.
func (s *CustomerService) Get(rs app.RequestScope, id int) (*models.Customer, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Customer.
func (s *CustomerService) Create(rs app.RequestScope, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Cust_id)
}

// Update updates the customer with the specified ID.
func (s *CustomerService) Update(rs app.RequestScope, id int, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the customer with the specified ID.
func (s *CustomerService) Delete(rs app.RequestScope, id int) (*models.Customer, error) {
	customer, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return customer, err
}

// Count returns the number of customer.
func (s *CustomerService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the customer with the specified offset and limit.
func (s *CustomerService) Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error) {
	return s.dao.Query(rs, offset, limit)
}