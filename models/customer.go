package models


import (
	"errors"
	"github.com/go-ozzo/ozzo-validation"
	// Import the Radix.v2 pool package, NOT the redis package.
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)

// // Declare a global db variable to store the Redis connection pool - already declared in "items.go"

func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    db, err = pool.New("tcp", "localhost:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}

// Create a new error message and store it as a constant. We'll use this
// error later if the FindCustomer() function fails to find an customer with a
// specific id.
var ErrNoCustomer = errors.New("models: no customers found")

type Customer struct {
	Cust_id int `json:"cust_id" db:"pk"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Post_address string `json:"post_address"`
}

 // Validate validates the Customer fields.
func (m Customer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.First_name, validation.Required, validation.Length(0, 50)),
	)
}

func PopulateCustomer(reply map[string]string) (*Customer, error) {
    var err error
    ab := new(Customer)
    ab.Cust_id, err = strconv.Atoi(reply["cust_id"])
    if err != nil {
        return nil, err
    }    
    ab.First_name = reply["first_name"]
    ab.Last_name = reply["last_name"]
    ab.Post_address = reply["postal_address"]
    return ab, nil
}

func FindCustomer(id string) (*Customer, error) {
    // Fetch the details of a specific c. If no customer is found with the
    // given id, the map[string]string returned by the Map() helper method
    // will be empty. So we can simply check whether it's length is zero and
    // return an ErrNoCustomer message if necessary.
    reply, err := db.Cmd("HGETALL", "cust_id:" + id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoCustomer
    }

    return PopulateCustomer(reply)
}

func CacheCustomer(customer *Customer) error {
    resp := db.Cmd("HMSET", "cust_id:" + strconv.Itoa(customer.Cust_id), "cust_id", strconv.Itoa(customer.Cust_id), "first_name", customer.First_name, "last_name", customer.Last_name, "post_address", customer.Post_address)
    if resp.Err != nil {
        log.Fatal(resp.Err)
        return resp.Err
    }
    return nil
}