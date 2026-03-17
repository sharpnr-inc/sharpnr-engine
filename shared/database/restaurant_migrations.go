package database

import (
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"gorm.io/gorm"
)

func migrateRestaurant(db *gorm.DB) error {
	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS hospitality`,
	).Error; err != nil {
		return err
	}
	return db.AutoMigrate(
		// =====================================================
		// Core
		// =====================================================
		&restaurant_models.Restaurant{},
		&restaurant_models.RestaurantFloor{},
		&restaurant_models.RestaurantTable{},

		// =====================================================
		// Settings
		// =====================================================
		&restaurant_models.RestaurantSettings{},
		&restaurant_models.RestaurantOrderNumberingSettings{},
		&restaurant_models.RestaurantTaxSettings{},
		&restaurant_models.RestaurantInvoiceSettings{},
		&restaurant_models.RestaurantKitchenSettings{},
		&restaurant_models.RestaurantPOSSettings{},
		&restaurant_models.RestaurantPreferedKitchen{},
		&restaurant_models.RestaurantConfiguration{},

		// =====================================================
		// Dining Flow
		// =====================================================
		&restaurant_models.RestaurantDiningSession{},

		// =====================================================
		// Orders
		// =====================================================
		&restaurant_models.RestaurantOrders{},
		&restaurant_models.RestaurantOrderItems{},
		&restaurant_models.RestaurantOrderItemsVariants{},
		&restaurant_models.RestaurantOrderItemsAddOns{},
		&restaurant_models.RestaurantPickUpOrders{},
	)
}
