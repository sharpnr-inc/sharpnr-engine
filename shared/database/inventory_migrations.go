package database

import (
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"gorm.io/gorm"
)

func migrateInventoryEngine(db *gorm.DB) error {
	db.AutoMigrate(
		&inventory_engine_models.InventoryAccount{},
		&inventory_engine_models.Inventory{},
		&inventory_engine_models.InventoryItem{},
		&inventory_engine_models.InventoryItemPrice{},
		&inventory_engine_models.InventoryItemHistory{},
	)
	return nil

}
