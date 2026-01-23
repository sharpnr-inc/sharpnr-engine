package database

import (
	"sharpnr-engine/core/organization/models"

	"gorm.io/gorm"
)

func migrateCore(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.Organization{},
		&models.OrganizationSubscription{},
		&models.OrganizationOffice{},
		&models.OrganizationSetting{},
		&models.OrganizationVerification{},
		&models.OrganizationSetupStatus{},
		&models.OrganizationGeoLocation{},
	)
}
