package models

type Promotion struct {
	Promo_id int `json:"promo_id" db:"pk"`
	Required_item_id int `json:"required_item_id"`
	Required_quantity int `json:"required_quantity"`
	Discount_percentage int `json:"discount_percentage"`
}