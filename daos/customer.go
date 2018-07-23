package daos

import (
	"../app"
	"../models"
)

// CustomerDAO persists data in database
type CustomerDAO struct{}

func NewCustomerDAO() *CustomerDAO {
	return &CustomerDAO{}
}

// Get reads the customer with the specified ID from the database.
func (dao *CustomerDAO) Get(rs app.RequestScope, id int) (*models.Customer, error) {
	var customer models.Customer
	err := rs.Tx().Select().Model(id, &customer)
	return &customer, err
}

// Create saves a new customer record in the database.
// The Customer.cust_id field will be populated with an automatically generated ID upon successful saving.
func (dao *CustomerDAO) Create(rs app.RequestScope, customer *models.Customer) error {
	customer.Cust_id = 0
	return rs.Tx().Model(customer).Insert()
}

// Update saves the changes to an customer in the database.
func (dao *CustomerDAO) Update(rs app.RequestScope, id int, customer *models.Customer) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	customer.Cust_id = id
	return rs.Tx().Model(customer).Exclude("Id").Update()
}

// Delete deletes an customer with the specified ID from the database.
func (dao *CustomerDAO) Delete(rs app.RequestScope, id int) error {
	customer, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(customer).Delete()
}

// Count returns the number of the customer records in the database.
func (dao *CustomerDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("customer").Row(&count)
	return count, err
}

// Query retrieves the customer records with the specified offset and limit from the database.
func (dao *CustomerDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Customer, error) {
	customer := []models.Customer{}
	err := rs.Tx().Select().OrderBy("cust_id").Offset(int64(offset)).Limit(int64(limit)).All(&customer)
	return customer, err
}

