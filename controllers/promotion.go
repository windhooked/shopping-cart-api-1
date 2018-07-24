package controllers

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"../app"
	"../models"
)

type (
	// promotionService specifies the interface for the artist service needed by artistResource.
	promotionService interface {
		Get(rs app.RequestScope, id int) (*models.Promotion, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Promotion, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Promotion) (*models.Promotion, error)
		Update(rs app.RequestScope, id int, model *models.Promotion) (*models.Promotion, error)
		Delete(rs app.RequestScope, id int) (*models.Promotion, error)
	}

	// promotionResource defines the handlers for the CRUD APIs.
	promotionResource struct {
		service promotionService
	}
)

func ServePromotionResource(rg *routing.RouteGroup, service promotionService) {
	r := &promotionResource{service}
	rg.Get("/promotions/<id>", r.get)
	rg.Get("/promotions", r.query)
	rg.Post("/promotions", r.create)
	rg.Put("/promotions/<id>", r.update)
	rg.Delete("/promotions/<id>", r.delete)
}

func (r *promotionResource) get(c *routing.Context) error {
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

func (r *promotionResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	promotions, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = promotions
	return c.Write(paginatedList)
}

func (r *promotionResource) create(c *routing.Context) error {
	var model models.Promotion
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *promotionResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *promotionResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

return c.Write(response)
}

