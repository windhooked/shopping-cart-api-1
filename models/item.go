package models


import (
	"errors"
	"github.com/go-ozzo/ozzo-validation"
	// Import the Radix.v2 pool package, NOT the redis package.
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)
// Declare a global db variable to store the Redis connection pool.
var db *pool.Pool

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
// error later if the FindItem() function fails to find an item with a
// specific id.
var ErrNoItem = errors.New("models: no items found")

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

func populateItem(reply map[string]string) (*Item, error) {
    var err error
    ab := new(Item)
    ab.Item_id, err = strconv.Atoi(reply["item_id"])
    if err != nil {
        return nil, err
    }    
    
    if reply["promo_id"] == "" {
        ab.Promo_id = nil
    } else {
        Promo_id_int_form, err := strconv.Atoi(reply["promo_id"])
        ab.Promo_id = &Promo_id_int_form
        
        if err != nil {
            return nil, err
        }    
    }
    
    ab.Name = reply["name"]
    ab.Stock, err = strconv.Atoi(reply["stock"])
    if err != nil {
        return nil, err
    }    
    ab.Price, err = strconv.Atoi(reply["price"])
    if err != nil {
        return nil, err
    } 
    return ab, nil
}

func FindItem(id string) (*Item, error) {
    // Fetch the details of a specific item. If no item is found with the
    // given id, the map[string]string returned by the Map() helper method
    // will be empty. So we can simply check whether it's length is zero and
    // return an ErrNoItem message if necessary.
    reply, err := db.Cmd("HGETALL", "item_id:" + id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoItem
    }

    return populateItem(reply)
}

// todo - implementing caching to redis
func CacheItem(item *Item) error {
    var converted_promo_id string
    if item.Promo_id == nil {
        converted_promo_id = ""
    } else {
        converted_promo_id = strconv.Itoa(*item.Promo_id)
    }
    resp := db.Cmd("HMSET", "item_id:" + strconv.Itoa(item.Item_id), "item_id", strconv.Itoa(item.Item_id), "promo_id", converted_promo_id, "name", item.Name, "stock", strconv.Itoa(item.Stock), "price", strconv.Itoa(item.Price))
    if resp.Err != nil {
        log.Fatal(resp.Err)
        return resp.Err
    }
    return nil
}