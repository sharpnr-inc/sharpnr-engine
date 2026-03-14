package legal_entity_repo

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgLegalEntityRepo) CreateTaxCategoryAccountMapping(ctx context.Context, mapping *legal_entity_models.TaxCategoryAccountMapping) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(mapping).Error
}

func (r *orgLegalEntityRepo) GetTaxCategoryByID(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_models.LegalEntityTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var taxCategory legal_entity_models.LegalEntityTaxCategory

	err := db.
		Where("id = ?", taxCategoryID).
		First(&taxCategory).
		Error

	if err != nil {
		return legal_entity_models.LegalEntityTaxCategory{}, err
	}

	return taxCategory, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetTaxCategoryAccountMappingByTaxCategoryID(
	ctx context.Context,
	LegalEntityTaxCategoryID uuid.UUID,
) (*legal_entity_models.TaxCategoryAccountMapping, error) {

	var mapping legal_entity_models.TaxCategoryAccountMapping

	err := r.db.WithContext(ctx).
		Where("legal_entity_tax_category_id = ? AND is_active = ?", LegalEntityTaxCategoryID, true).
		First(&mapping).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &mapping, nil
}
