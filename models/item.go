package models


import "github.com/go-ozzo/ozzo-validation"

// Item represents an item record in DB
type Item struct {
	Item_id int
	Promo_id int
	Name string
	Stock int
	Price int
}

 // Validate validates the Item fields.
func (m Item) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 50)),
	)
}