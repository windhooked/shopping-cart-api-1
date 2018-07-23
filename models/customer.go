package models


import "github.com/go-ozzo/ozzo-validation"

type Customer struct {
	Cust_id int `json:"cust_id" db:"pk"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Post_address string `json:"post_address"`
}

 // Validate validates the Item fields.
func (m Customer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.First_name, validation.Required, validation.Length(0, 50)),
	)
}