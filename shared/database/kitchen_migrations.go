package database

import (
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"gorm.io/gorm"
)

func migrateKitchen(db *gorm.DB) error {
	return db.AutoMigrate(
		//Hospitality --> Kitchen
		&kitchen_models.Kitchen{},
		&kitchen_models.KitchenTicket{},
		&kitchen_models.KitchenTicketItem{},
		&kitchen_models.KitchenTicketItemAddon{},
	)
}
