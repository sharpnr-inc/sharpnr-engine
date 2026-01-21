package database

import (
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"gorm.io/gorm"
)

func migrateHospitalityMenu(db *gorm.DB) error {

	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS hospitality`,
	).Error; err != nil {
		return err
	}
	return db.AutoMigrate(

		&hospitalityMenu_models.MenuGroup{},
		&hospitalityMenu_models.MenuCategory{},
		&hospitalityMenu_models.MenuSubCategory{},
		&hospitalityMenu_models.MenuItem{},
		&hospitalityMenu_models.MenuVariant{},
		&hospitalityMenu_models.MenuVariantGroup{},
		&hospitalityMenu_models.MenuItemAddon{},
	)
}
