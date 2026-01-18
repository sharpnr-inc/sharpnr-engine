package database

import (
	org_property_models "sharpnr-engine/core/org_property/models"

	"gorm.io/gorm"
)

func migrateOrgProperty(db *gorm.DB) error {
	return db.AutoMigrate(
		&org_property_models.OrganizationProperty{},
		&org_property_models.OrganizationPropertyAddress{},
		&org_property_models.OrganizationPropertyDomain{},
		&org_property_models.OrganizationPropertyTimeZone{},
	)
}
