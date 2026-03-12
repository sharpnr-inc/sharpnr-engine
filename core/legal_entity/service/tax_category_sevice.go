package legal_entity_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) AddLegalEntityTaxCategory(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	request legal_entity_contracts.CreateLegalEntityTaxCategory,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		taxCategory := &legal_entity_models.LegalEntityTaxCategory{
			ID:            uuid.New(),
			LegalEntityID: request.LegalEntityID,
			Name:          request.CategoryName,
			Code:          request.Code,
			Description:   request.Description,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := s.legalEntityRepo.CreateTaxCategory(txCtx, taxCategory); err != nil {
			return err
		}

		taxCategoryRule := &legal_entity_models.LegalEntityTaxRule{
			ID:                           uuid.New(),
			LegalEntityID:                request.LegalEntityID,
			LegalEntityTaxCategoryID:     taxCategory.ID,
			LegalEntityTaxRegistrationID: request.Rule.LegalEntityTaxRegistrationID,
			TaxType:                      "GST/VAT",
			Rate:                         request.Rule.Rate,
			MinValue:                     request.Rule.MinValue,
			MaxValue:                     request.Rule.MaxValue,
			EffectiveFrom:                request.Rule.EffectiveFrom,
			EffectiveTo:                  request.Rule.EffectiveTo,
			IsActive:                     true,
		}

		if err := s.legalEntityRepo.CreateTaxRuleOfTaxCategory(txCtx, taxCategoryRule); err != nil {
			return err
		}

		for _, tc := range request.Rule.Components {

			taxRuleComponent := &legal_entity_models.LegalEntityTaxComponent{
				ID:                   uuid.New(),
				LegalEntityTaxRuleID: taxCategoryRule.ID,
				ComponentCode:        tc.ComponentCode,
				Rate:                 tc.Rate,
				LiabilityAccountID:   tc.LiabilityAccountID,
				CreatedAt:            time.Now(),
			}

			if err := s.legalEntityRepo.CreateTaxComponentOfTaxRule(txCtx, taxRuleComponent); err != nil {
				return err
			}
		}

		return nil

	})
}
