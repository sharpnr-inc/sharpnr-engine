package database

import (
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"gorm.io/gorm"
)

func migrateOrgLegalEntity(db *gorm.DB) error {

	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS legal_entity`,
	).Error; err != nil {
		return err
	}

	return db.AutoMigrate(

		&org_legal_entity_models.OrganizationLegalEntity{},
		&org_legal_entity_models.PropertyLegalEntity{},
		&org_legal_entity_models.LegalEntityAddress{},
		&org_legal_entity_models.LegalEntityTaxRegistration{},
		&org_legal_entity_models.ModuleLegalEntity{},
		&org_legal_entity_models.LegalEntityCompliance{},
		&org_legal_entity_models.LegalEntityAccountingMapping{},
		&org_legal_entity_models.LegalEntityTaxCategory{},
		&org_legal_entity_models.LegalEntityTaxRule{},
	)
}
