package models

type Order struct {
	Order_id int `json:"order_id" db:"pk"`
	Cust_id int `json:"cust_id"`
	Item_id int `json:"item_id"`
	Quantity int `json:"quantity"`
	Placed bool `json:"placed"`
}