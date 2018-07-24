package services

import (
	"../app"
	"../models"
)

// promotionDAO specifies the interface of the promotion DAO needed by PromotionService.
type promotionDAO interface {
	// Get returns the promotion with the specified promotion ID.
	Get(rs app.RequestScope, id int) (*models.Promotion, error)
	// Count returns the number of promotions.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of promotions with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Promotion, error)
	// Create saves a new promotion in the storage.
	Create(rs app.RequestScope, promotion *models.Promotion) error
	// Update updates the promotion with given ID in the storage.
	Update(rs app.RequestScope, id int, promotion *models.Promotion) error
	// Delete removes the promotion with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// PromotionService provides services related with promotions.
type PromotionService struct {
	dao promotionDAO
}

// NewPromotionService creates a new PromotionService with the given promotion DAO.
func NewPromotionService(dao promotionDAO) *PromotionService {
	return &PromotionService{dao}
}

// Get returns the promotion with the specified the promotion ID.
func (s *PromotionService) Get(rs app.RequestScope, id int) (*models.Promotion, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Promotion.
func (s *PromotionService) Create(rs app.RequestScope, model *models.Promotion) (*models.Promotion, error) {
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Promo_id)
}

// Update updates the promotion with the specified ID.
func (s *PromotionService) Update(rs app.RequestScope, id int, model *models.Promotion) (*models.Promotion, error) {
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the promotion with the specified ID.
func (s *PromotionService) Delete(rs app.RequestScope, id int) (*models.Promotion, error) {
	promotion, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return promotion, err
}

// Count returns the number of promotions.
func (s *PromotionService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the promotions with the specified offset and limit.
func (s *PromotionService) Query(rs app.RequestScope, offset, limit int) ([]models.Promotion, error) {
	return s.dao.Query(rs, offset, limit)
}