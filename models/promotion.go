package models

import (
	"errors"
	// Import the Radix.v2 pool package, NOT the redis package.
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)

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
// error later if the FindPromotion() function fails to find an promotion with a
// specific id.
var ErrNoPromotion = errors.New("models: no promotions found")

type Promotion struct {
	Promo_id int `json:"promo_id" db:"pk"`
	Required_item_id int `json:"required_item_id"`
	Required_quantity int `json:"required_quantity"`
	Discount_percentage int `json:"discount_percentage"`
}

func PopulatePromotion(reply map[string]string) (*Promotion, error) {
    var err error
    ab := new(Promotion)
    ab.Promo_id, err = strconv.Atoi(reply["promo_id"])
    if err != nil {
        return nil, err
    }    
    ab.Required_item_id, err = strconv.Atoi(reply["required_item_id"])
    if err != nil {
        return nil, err
    }    
    ab.Required_quantity, err = strconv.Atoi(reply["required_quantity"])
    if err != nil {
        return nil, err
    } 
	ab.Discount_percentage, err = strconv.Atoi(reply["discount_percentage"])
    if err != nil {
        return nil, err
    }     
    return ab, nil
}

func FindPromotion(id string) (*Promotion, error) {
    // Fetch the details of a specific promotion. If no promotion is found with the
    // given id, the map[string]string returned by the Map() helper method
    // will be empty. So we can simply check whether it's length is zero and
    // return an ErrNoPromotion message if necessary.
    reply, err := db.Cmd("HGETALL", "promo_id:" + id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoPromotion
    }

    return PopulatePromotion(reply)
}

func CachePromotion(promotion *Promotion) error {
    resp := db.Cmd("HMSET", "promo_id:" + strconv.Itoa(promotion.Promo_id), "promo_id", strconv.Itoa(promotion.Promo_id), "required_item_id", strconv.Itoa(promotion.Required_item_id), "required_quantity", strconv.Itoa(promotion.Required_quantity), "discount_percentage", strconv.Itoa(promotion.Discount_percentage))
    if resp.Err != nil {
        log.Fatal(resp.Err)
        return resp.Err
    }
    return nil
}