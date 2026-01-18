package database

import (
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"gorm.io/gorm"
)

func migrateOrgLegalEntity(db *gorm.DB) error {
	return db.AutoMigrate(
		&org_legal_entity_models.OrganizationLegalEntity{},
		&org_legal_entity_models.PropertyLegalEntity{},
		&org_legal_entity_models.LegalEntityTaxCategory{},
		&org_legal_entity_models.LegalEntityAddress{},
		&org_legal_entity_models.LegalEntityTaxRegistration{},
		&org_legal_entity_models.ModuleLegalEntity{},
		&org_legal_entity_models.LegalEntityCompliance{},
	)
}
