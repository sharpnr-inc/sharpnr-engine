package database

import (
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"gorm.io/gorm"
)

func migrateRestaurant(db *gorm.DB) error {
	return db.AutoMigrate(
		&restaurant_models.Restaurant{},
		&restaurant_models.RestaurantFloor{},
		&restaurant_models.RestaurantTable{},
		&restaurant_models.RestaurantDiningSession{},
		&restaurant_models.RestaurantCharge{},
		&restaurant_models.RestaurantSessionPayment{},
	)
}
