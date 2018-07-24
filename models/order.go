package models

type Purchase_Order struct {
	Purchase_order_id int `json:"purchase_order_id" db:"pk"`
	Cust_id int `json:"cust_id"`
	Item_id int `json:"item_id"`
	Quantity int `json:"quantity"`
	Placed bool `json:"placed"`
}