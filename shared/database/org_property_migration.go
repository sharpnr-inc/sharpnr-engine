package database

import (
	org_property_models "sharpnr-engine/core/org_property/models"

	"gorm.io/gorm"
)

func migrateOrgProperty(db *gorm.DB) error {
	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS organization_property`,
	).Error; err != nil {
		return err
	}
	return db.AutoMigrate(
		&org_property_models.OrganizationProperty{},
		&org_property_models.OrganizationPropertyAddress{},
		&org_property_models.OrganizationPropertyDomain{},
		&org_property_models.OrganizationPropertyTimeZone{},
		&org_property_models.OrganizationPropertyGeoLocation{},
		&org_property_models.PropertyServicePoint{},
		&org_property_models.ServicePointKitchenMapping{},
	)
}
