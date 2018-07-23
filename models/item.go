package models


import "github.com/go-ozzo/ozzo-validation"

// Item represents an item record in DB
// Id   int    `json:"id" db:"id"`
// Name string `json:"name" db:"name"`
type Item struct {
	Item_id int `json:"item_id" db:"pk"`
	Promo_id *int `json:"promo_id"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	Price int `json:"price"`
}

 // Validate validates the Item fields.
func (m Item) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 50)),
	)
}