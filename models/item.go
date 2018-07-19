package models


import "github.com/go-ozzo/ozzo-validation"

// Item represents an item record in DB
type Item struct {
	item_id int
	promo_id int
	name string
	stock int
	price int
}

 // Validate validates the Item fields.
func (m Item) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.name, validation.Required, validation.Length(0, 50)),
	)
}

// CREATE TABLE Item
// (
//     item_id SERIAL PRIMARY KEY,
//     promo_id INTEGER,
//     name VARCHAR(50) NOT NULL,
//     stock INTEGER NOT NULL,
//     price INTEGER NOT NULL
// );