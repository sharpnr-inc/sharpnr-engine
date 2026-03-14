package database

import (
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"gorm.io/gorm"
)

func migrateLegalEntity(db *gorm.DB) error {

	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS legal_entity`,
	).Error; err != nil {
		return err
	}

	return db.AutoMigrate(

		&legal_entity_models.LegalEntity{},
		&legal_entity_models.PropertyLegalEntity{},
		&legal_entity_models.LegalEntityAddress{},
		&legal_entity_models.LegalEntityTaxRegistration{},
		&legal_entity_models.ModuleLegalEntity{},
		&legal_entity_models.LegalEntityCompliance{},
		&legal_entity_models.TaxCategoryAccountMapping{},
		&legal_entity_models.LegalEntityTaxCategory{},
		&legal_entity_models.LegalEntityTaxRule{},
		&legal_entity_models.LegalEntityTaxComponent{},
	)
}
