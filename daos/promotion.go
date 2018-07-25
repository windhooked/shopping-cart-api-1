package daos

import (
	"../app"
	"../models"
	"strconv"
)

// PromotionDAO persists data in database
type PromotionDAO struct{}

func NewPromotionDAO() *PromotionDAO {
	return &PromotionDAO{}
}

// Get reads the promotion with the specified ID from the database.
func (dao *PromotionDAO) Get(rs app.RequestScope, id int) (*models.Promotion, error) {
	// Check if order is cached in Redis otherwise use query to the DB
    cachedPromotion, err := models.FindPromotion(strconv.Itoa(id))
    if err == models.ErrNoPromotion {
		var promotion models.Promotion
		err := rs.Tx().Select().Model(id, &promotion)
		models.CachePromotion(&promotion)
		return &promotion, err
    } else if err != nil {
        return nil, err
    } else {
    	return cachedPromotion, err
    }	
}

// Create saves a new promotion record in the database.
// The Promotion.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *PromotionDAO) Create(rs app.RequestScope, promotion *models.Promotion) error {
	promotion.Promo_id = 0
	return rs.Tx().Model(promotion).Insert()
}

// Update saves the changes to an promotion in the database.
func (dao *PromotionDAO) Update(rs app.RequestScope, id int, promotion *models.Promotion) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	promotion.Promo_id = id
	return rs.Tx().Model(promotion).Exclude("Id").Update()
}

// Delete deletes an promotion with the specified ID from the database.
func (dao *PromotionDAO) Delete(rs app.RequestScope, id int) error {
	promotion, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(promotion).Delete()
}

// Count returns the number of the promotion records in the database.
func (dao *PromotionDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("promotion").Row(&count)
	return count, err
}

// Query retrieves the promotion records with the specified offset and limit from the database.
func (dao *PromotionDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Promotion, error) {
	promotion := []models.Promotion{}
	err := rs.Tx().Select().OrderBy("promo_id").Offset(int64(offset)).Limit(int64(limit)).All(&promotion)
	return promotion, err
}