package restaurantrepo

import "gorm.io/gorm"

type restaurantRepo struct {
	db *gorm.DB
}

// CreateRestaurant implements RestaurantRepository.
func (r *restaurantRepo) CreateRestaurant() {
	panic("unimplemented")
}

func NewRestaurantRepository(db *gorm.DB) *restaurantRepo {
	return &restaurantRepo{
		db: db,
	}
}
