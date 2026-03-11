package legal_entity_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) CreateTaxRegistration(ctx context.Context, reg *legal_entity_models.LegalEntityTaxRegistration) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(reg).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetTaxRegistrations(ctx context.Context, legalentityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRegistration, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var registrations []legal_entity_models.LegalEntityTaxRegistration

	// Use Where to filter by the Legal Entity and Order by date descending
	err := db.WithContext(ctx).
		Where("organization_legal_entity_id = ?", legalentityID).
		Order("created_at DESC").
		Find(&registrations).
		Error

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

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
		Where("organization_legal_entity_id = ?", entityID).
		Order("code ASC").
		Find(&categories).
		Error

	return categories, err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetTaxCategoryByCode(ctx context.Context, entityID uuid.UUID, code string) (*legal_entity_models.LegalEntityTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var category legal_entity_models.LegalEntityTaxCategory

	err := db.WithContext(ctx).
		Where("organization_legal_entity_id = ? AND code = ?", entityID, code).
		First(&category).
		Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

/* ---------------------------------- <***> --------------------------------- */
