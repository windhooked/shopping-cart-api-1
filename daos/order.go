package daos

import (
	"../app"
	"../models"
)

// OrderDAO persists data in database
type OrderDAO struct{}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{}
}

// Get reads the order with the specified ID from the database.
func (dao *OrderDAO) Get(rs app.RequestScope, id int) (*models.Order, error) {
	var order models.Order
	err := rs.Tx().Select().Model(id, &order)
	return &order, err
}

// Create saves a new order record in the database.
// The Order.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *OrderDAO) Create(rs app.RequestScope, order *models.Order) error {
	order.Order_id = 0
	return rs.Tx().Model(order).Insert()
}

// Update saves the changes to an order in the database.
func (dao *OrderDAO) Update(rs app.RequestScope, id int, order *models.Order) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	order.Order_id = id
	return rs.Tx().Model(order).Exclude("Id").Update()
}

// Delete deletes an order with the specified ID from the database.
func (dao *OrderDAO) Delete(rs app.RequestScope, id int) error {
	order, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(order).Delete()
}

// Count returns the number of the order records in the database.
func (dao *OrderDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("Order").Row(&count)
	return count, err
}

// Query retrieves the order records with the specified offset and limit from the database.
func (dao *OrderDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Order, error) {
	order := []models.Order{}
	err := rs.Tx().Select().From("Order").OrderBy("order_id").Offset(int64(offset)).Limit(int64(limit)).All(&order)
	return order, err
}

