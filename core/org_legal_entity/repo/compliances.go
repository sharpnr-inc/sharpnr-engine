package org_legal_entity_repo

import (
	"context"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"github.com/google/uuid"
)

func (o *orgLegalEntityRepo) CreateLegalEntityCompliance(ctx context.Context, compliance *org_legal_entity_models.LegalEntityCompliance) error {
	// .Create adds the compliance record (e.g., a specific permit or tax status)
	return o.db.WithContext(ctx).Create(compliance).Error
}

func (o *orgLegalEntityRepo) GetLegalCompliancesByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]*org_legal_entity_models.LegalEntityCompliance, error) {
	var compliances []*org_legal_entity_models.LegalEntityCompliance

	// We filter by organization_legal_entity_id.
	// Ordering by 'expiry_date' or 'created_at' is usually helpful for compliance.
	err := o.db.WithContext(ctx).
		Where("organization_legal_entity_id = ?", legalEntityID).
		Order("created_at DESC").
		Find(&compliances).Error

	if err != nil {
		return nil, err
	}

	return compliances, nil
}
