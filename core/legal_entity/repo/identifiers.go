package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *orgLegalEntityRepo) AddLegalEntityIdentifier(
	ctx context.Context,
	req legal_entity_models.LegalEntityIdentifier,
) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	if err := db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}

	return nil
}

func (d *orgLegalEntityRepo) FindAllIdentitifiersOfLegalEntity(
	ctx context.Context,
	legalEntityID uuid.UUID,
) ([]legal_entity_models.LegalEntityIdentifier, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	var identifiers []legal_entity_models.LegalEntityIdentifier

	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&identifiers).
		Error

	if err != nil {
		return nil, err
	}

	// Explicitly return gorm.ErrRecordNotFound if no identifiers are found
	if len(identifiers) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return identifiers, nil
}

func (d *orgLegalEntityRepo) GetIdentifiersByID(
	ctx context.Context,
	identifierID uuid.UUID,
) (*legal_entity_models.LegalEntityIdentifier, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	var identifier legal_entity_models.LegalEntityIdentifier

	// First automatically returns gorm.ErrRecordNotFound if the ID doesn't exist
	err := db.WithContext(ctx).
		First(&identifier, "id = ?", identifierID).
		Error

	if err != nil {
		return nil, err
	}

	return &identifier, nil
}
