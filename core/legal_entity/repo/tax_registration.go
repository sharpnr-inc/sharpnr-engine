package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *orgLegalEntityRepo) AddNewTaxRegistration(ctx context.Context, reg *legal_entity_models.LegalEntityTaxRegistration) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(reg).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) FindAllTaxRegistrationOfLegalEntity(ctx context.Context, legalentityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRegistration, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var registrations []legal_entity_models.LegalEntityTaxRegistration

	// Use Where to filter by the Legal Entity and Order by date descending
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalentityID).
		Order("created_at DESC").
		Find(&registrations).
		Error

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

/* ---------------------------------- <***> --------------------------------- */
