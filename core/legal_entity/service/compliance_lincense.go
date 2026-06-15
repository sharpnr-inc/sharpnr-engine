package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) AddNewLicenseOrComplianceInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req legal_entity_contracts.AddNewLicenseAndComplianceRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.legalEntityRepo.CreateLegalEntityCompliance(txCtx, &legal_entity_models.LegalEntityCompliance{
			ID:               uuid.New(),
			LegalEntityID:    legalEntityID,
			Type:             req.Type,
			LicenseNumber:    req.LicenseNumber,
			IssuingAuthority: req.IssuingAuthority,
			ValidFrom:        req.ValidFrom,
			ValidTo:          req.ValidTo,
			IsActive:         false,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}); err != nil {
			return nil
		}

		return nil
	})
}
