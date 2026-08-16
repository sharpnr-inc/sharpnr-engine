package restaurantrepo

import (
	"gorm.io/gorm"
)

type restaurantRepo struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *restaurantRepo {
	return &restaurantRepo{
		db: db,
	}
}
