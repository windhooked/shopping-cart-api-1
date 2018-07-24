package services

import (
	// "fmt"
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

// // itemDAO specifies the interface of the item DAO needed by ItemService.
// type itemDAO interface {
// 	// Get returns the item with the specified item ID.
// 	Get(rs app.RequestScope, id int) (*models.Item, error)
// 	// Count returns the number of items.
// 	Count(rs app.RequestScope) (int, error)
// 	// Query returns the list of items with the given offset and limit.
// 	Query(rs app.RequestScope, offset, limit int) ([]models.Item, error)
// 	// Create saves a new item in the storage.
// 	Create(rs app.RequestScope, item *models.Item) error
// 	// Update updates the item with given ID in the storage.
// 	Update(rs app.RequestScope, id int, item *models.Item) error
// 	// Delete removes the item with given ID from the storage.
// 	Delete(rs app.RequestScope, id int) error
// }

// // orderDAO specifies the interface of the order DAO needed by OrderService.
// type orderDAO interface {
// 	// Get returns the order with the specified order ID.
// 	Get(rs app.RequestScope, id int) (*models.Purchase_Order, error)
// 	// Count returns the number of orders.
// 	Count(rs app.RequestScope) (int, error)
// 	// Query returns the list of orders with the given offset and limit.
// 	Query(rs app.RequestScope, offset, limit int) ([]models.Purchase_Order, error)
// 	// Create saves a new order in the storage.
// 	Create(rs app.RequestScope, order *models.Purchase_Order) error
// 	// Update updates the order with given ID in the storage.
// 	Update(rs app.RequestScope, id int, order *models.Purchase_Order) error
// 	// Delete removes the order with given ID from the storage.
// 	Delete(rs app.RequestScope, id int) error
// }

// // promotionDAO specifies the interface of the promotion DAO needed by PromotionService.
// type promotionDAO interface {
// 	// Get returns the promotion with the specified promotion ID.
// 	Get(rs app.RequestScope, id int) (*models.Promotion, error)
// 	// Count returns the number of promotions.
// 	Count(rs app.RequestScope) (int, error)
// 	// Query returns the list of promotions with the given offset and limit.
// 	Query(rs app.RequestScope, offset, limit int) ([]models.Promotion, error)
// 	// Create saves a new promotion in the storage.
// 	Create(rs app.RequestScope, promotion *models.Promotion) error
// 	// Update updates the promotion with given ID in the storage.
// 	Update(rs app.RequestScope, id int, promotion *models.Promotion) error
// 	// Delete removes the promotion with given ID from the storage.
// 	Delete(rs app.RequestScope, id int) error
// }

// CustomerService provides services related with customers.
type CustomerService struct {
	cust_dao customerDAO
	item_dao itemDAO
	order_dao orderDAO
	promotion_dao promotionDAO
}

// NewCustomerService creates a new CustomerService with the given customer DAO.
func NewCustomerService(cust_dao customerDAO, item_dao itemDAO, order_dao orderDAO, promotion_dao promotionDAO) *CustomerService {
	return &CustomerService{cust_dao, item_dao, order_dao, promotion_dao}
}

// Get returns the customer with the specified the cust ID.
func (s *CustomerService) Get(rs app.RequestScope, id int) (*models.Customer, error) {
	return s.cust_dao.Get(rs, id)
}

// Returns the current shopping cart for the customer with the specified cust ID
func (s *CustomerService) GetCart(rs app.RequestScope, id int) (string, error) {
	customer, err := s.cust_dao.Get(rs, id)
	if err != nil {
		return "nil", err
	}
	println(customer.Cust_id)
	// println(customer.First_name)
	// println(customer.Last_name)
	// println(customer.Post_address)

	// item, err := s.item_dao.Get(rs, id)
	// if err != nil {
	// 	return "nil", err
	// }
	order, err := s.order_dao.GetByCustomerId(rs, id)
	if err != nil {
		return "nil", err
	}
	print(order.Cust_id)

	// promotion, err := s.promotion_dao.Get(rs, id)
	// if err != nil {
	// 	return "nil", err
	// }

	// fmt.Printf("%+v\n", customer)
	return "string", err
}

// Create creates a new Customer.
func (s *CustomerService) Create(rs app.RequestScope, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.cust_dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.cust_dao.Get(rs, model.Cust_id)
}

// Update updates the customer with the specified ID.
func (s *CustomerService) Update(rs app.RequestScope, id int, model *models.Customer) (*models.Customer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.cust_dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.cust_dao.Get(rs, id)
}

// Delete deletes the customer with the specified ID.
func (s *CustomerService) Delete(rs app.RequestScope, id int) (*models.Customer, error) {
	customer, err := s.cust_dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.cust_dao.Delete(rs, id)
	return customer, err
}

// Count returns the number of customer.
func (s *CustomerService) Count(rs app.RequestScope) (int, error) {
	return s.cust_dao.Count(rs)
}

// Query returns the customer with the specified offset and limit.
func (s *CustomerService) Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error) {
	return s.cust_dao.Query(rs, offset, limit)
}