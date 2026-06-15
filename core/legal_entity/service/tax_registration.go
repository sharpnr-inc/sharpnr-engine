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

func (s *orgLegalEntity) AddNewTaxRegistration(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req legal_entity_contracts.AddNewTaxRegistrationInLegalEntityRequestDto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.legalEntityRepo.AddNewTaxRegistration(txCtx, &legal_entity_models.LegalEntityTaxRegistration{

			ID:                  uuid.New(),
			LegalEntityID:       legalEntityID,
			TaxType:             req.TaxType,
			RegistrationNumber:  req.RegistrationNumber,
			GSTRegistrationType: req.GSTRegistrationType,
			State:               req.State,
			GSTStateCode:        req.GSTStateCode,
			Country:             req.Country,
			EffectiveFrom:       req.EffectiveFrom,
			EffectiveTo:         req.EffectiveTo,
			IsActive:            false,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}); err != nil {
			return err
		}
		return nil
	})
}
