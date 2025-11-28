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
