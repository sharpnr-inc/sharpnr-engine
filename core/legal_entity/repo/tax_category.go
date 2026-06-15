package legal_entity_repo

import (
	"context"
	"errors"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) CreateTaxCategory(ctx context.Context, category *legal_entity_models.LegalEntityTaxCategory) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	if category.ID == uuid.Nil {
		category.ID = uuid.New()
	}

	return db.WithContext(ctx).Create(category).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) ListTaxCategories(ctx context.Context, entityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxCategory, error) {
    db, ok := dbtx.ExtractTx(ctx)
    if !ok {
        db = r.db
    }

    var categories []legal_entity_models.LegalEntityTaxCategory

    err := db.WithContext(ctx).
        Where("legal_entity_id = ?", entityID).
        Order("code ASC").
        Find(&categories).
        Error
    if err != nil {
        return nil, err
    }

    // Since Find doesn't return ErrRecordNotFound, check the slice length
    if len(categories) == 0 {
        // You can return a custom error, or gorm.ErrRecordNotFound if you want to fake it
        return nil, gorm.ErrRecordNotFound 
    }

    return categories, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetTaxCategoryByCode(ctx context.Context, entityID uuid.UUID, code string) (*legal_entity_models.LegalEntityTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var category legal_entity_models.LegalEntityTaxCategory

	err := db.WithContext(ctx).
		Where("legal_entity_id = ? AND code = ?", entityID, code).
		First(&category).
		Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) CreateTaxCategoryAccountMapping(ctx context.Context, mapping *legal_entity_models.TaxCategoryAccountMapping) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(mapping).Error
}

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */
