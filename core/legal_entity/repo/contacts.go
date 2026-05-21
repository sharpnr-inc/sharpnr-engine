package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *orgLegalEntityRepo) AddLegalEntityContacts(
	ctx context.Context,
	req legal_entity_models.LegalEntityContact,
) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	// Persist the legal entity contact to the database
	if err := db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}

	return nil
}

func (d *orgLegalEntityRepo) FindAddAllContactsOfLegalEntity(
	ctx context.Context,
	legalEntityID uuid.UUID,
) ([]legal_entity_models.LegalEntityContact, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	var contacts []legal_entity_models.LegalEntityContact

	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&contacts).
		Error

	if err != nil {
		return nil, err
	}

	// Explicitly return gorm.ErrRecordNotFound if no records are matched
	if len(contacts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return contacts, nil
}
