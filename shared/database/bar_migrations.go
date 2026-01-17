package database

import (
	bar_models "sharpnr-engine/modules/hospitality/bar/models"

	"gorm.io/gorm"
)

func migrateBar(db *gorm.DB) error {
	return db.AutoMigrate(
		// Hospitality --> Bar
		&bar_models.Bar{},
		&bar_models.BarFloor{},
		&bar_models.BarTable{},
		// Bar-orders
		&bar_models.BarTableSession{},
		&bar_models.BarTablesSessionOrder{},
	)
}
