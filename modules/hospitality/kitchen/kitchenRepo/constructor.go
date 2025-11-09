package kitchenrepo

import "gorm.io/gorm"

type kitchenRepository struct {
	db *gorm.DB
}

func NewKitchenRepository(db *gorm.DB) *kitchenRepository {
	return &kitchenRepository{
		db: db,
	}
}

func (r *kitchenRepository) GetOrders() ([]interface{}, error) {
	// Stub – replace with actual GORM query (HospitalityOrder)
	return []interface{}{}, nil
}

func (r *kitchenRepository) UpdateOrder(orderID string, status string) error {
	// Stub – replace with GORM update logic
	return nil
}
