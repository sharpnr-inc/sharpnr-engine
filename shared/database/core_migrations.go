package database

import (
	"sharpnr-engine/core/organization/models"

	"gorm.io/gorm"
)

func migrateCore(db *gorm.DB) error {
	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS core`,
	).Error; err != nil {
		return err
	}

	return db.AutoMigrate(
		&models.Organization{},
		&models.OrganizationSubscription{},
		&models.OrganizationOffice{},
		&models.OrganizationSetting{},
		&models.OrganizationVerification{},
	)
}
