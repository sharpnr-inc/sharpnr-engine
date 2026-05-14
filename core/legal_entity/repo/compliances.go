package legal_entity_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

func (r *orgLegalEntityRepo) CreateLegalEntityCompliance(ctx context.Context, compliance *legal_entity_models.LegalEntityCompliance) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(compliance).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetLegalCompliancesByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]*legal_entity_models.LegalEntityCompliance, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var compliances []*legal_entity_models.LegalEntityCompliance

	// We filter by organization_legal_entity_id.
	// Ordering by 'expiry_date' or 'created_at' is usually helpful for compliance.
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Order("created_at DESC").
		Find(&compliances).Error

	if err != nil {
		return nil, err
	}

	return compliances, nil
}

/* ---------------------------------- <***> --------------------------------- */
