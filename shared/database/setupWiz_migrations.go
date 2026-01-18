package database

import (
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"

	"gorm.io/gorm"
)

func migrateSetupWizard(db *gorm.DB) error {
	return db.AutoMigrate(
		&org_setup_wizard_models.OrganizationSetup{},
		// &org_setup_wizard_models.OrganizationSetupStep{},
		// &org_setup_wizard_models.OrganizationSetupDraft{},
	)
}
