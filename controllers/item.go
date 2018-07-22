package controllers

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"../app"
	"../models"
)

type (
	// itemService specifies the interface for the artist service needed by artistResource.
	itemService interface {
		Get(rs app.RequestScope, id int) (*models.Item, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Item, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Item) (*models.Item, error)
		Update(rs app.RequestScope, id int, model *models.Item) (*models.Item, error)
		Delete(rs app.RequestScope, id int) (*models.Item, error)
	}

	// itemResource defines the handlers for the CRUD APIs.
	itemResource struct {
		service itemService
	}
)

func ServeItemResource(rg *routing.RouteGroup, service itemService) {
	r := &itemResource{service}
	rg.Get("/items/<id>", r.get)
	// rg.Get("/items", r.query)
	// rg.Post("/items", r.create)
	// rg.Put("/items/<id>", r.update)
	// rg.Delete("/items/<id>", r.delete)
}

func (r *itemResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

// func (r *itemResource) query(c *routing.Context) error {
// 	return
// }

// func (r *itemResource) create(c *routing.Context) error {
// 	return
// }

// func (r *itemResource) update(c *routing.Context) error {
// 	return
// }

// func (r *itemResource) delete(c *routing.Context) error {
// 	return
// }

