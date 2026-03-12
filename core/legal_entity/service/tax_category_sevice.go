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

		taxCategoryMapping := &legal_entity_models.TaxCategoryAccountMapping{
			ID:                       uuid.New(),
			LegalEntityID:            legalEntityID,
			LegalEntityTaxCategoryID: taxCategory.ID,
			AccountingBookID:         request.AccountingBookID,
			ReceivableAccountID:      request.ReceivableAccountID,
			IncomeAccountID:          request.IncomeAccountID,
			IsActive:                 true,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		}

		if err := s.legalEntityRepo.CreateTaxCategoryAccountMapping(txCtx, taxCategoryMapping); err != nil {
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

func (c *orgLegalEntity) GetAllLegalEntityTaxCategory(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error) {

	var res []legal_entity_contracts.LegalEntityTaxCategoryResponseDto

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		taxCategories, err := c.legalEntityRepo.ListTaxCategories(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		for _, tc := range taxCategories {

			// account mapping
			accountMapping, err := c.legalEntityRepo.GetTaxCategoryAccountMappingByTaxCategoryID(txCtx, tc.ID)
			if err != nil {
				return err
			}

			// rules
			rules, err := c.legalEntityRepo.GetActiveRulesByCategory(txCtx, tc.ID)
			if err != nil {
				return err
			}

			ruleDtos, err := c.buildRuleDtos(txCtx, rules)
			if err != nil {
				return err
			}

			categoryDto := legal_entity_contracts.LegalEntityTaxCategoryResponseDto{
				ID:           tc.ID,
				CategoryName: tc.Name,
				Code:         tc.Code,
				Description:  tc.Description,
				Rules:        ruleDtos,
			}

			if accountMapping != nil {
				categoryDto.AccountingBookID = accountMapping.AccountingBookID
				categoryDto.ReceivableAccountID = accountMapping.ReceivableAccountID
				categoryDto.IncomeAccountID = accountMapping.IncomeAccountID
			}

			res = append(res, categoryDto)
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *orgLegalEntity) buildRuleDtos(
	ctx context.Context,
	rules []legal_entity_models.LegalEntityTaxRule,
) ([]legal_entity_contracts.LegalEntityTaxCategoryTaxRuleResponseDto, error) {

	var ruleDtos []legal_entity_contracts.LegalEntityTaxCategoryTaxRuleResponseDto

	for _, rule := range rules {

		components, err := c.legalEntityRepo.GetTaxRuleComponents(ctx, rule.ID)
		if err != nil {
			return nil, err
		}

		componentDtos := make([]legal_entity_contracts.TaxComponentResponseDto, 0, len(components))

		for _, comp := range components {
			componentDtos = append(componentDtos, legal_entity_contracts.TaxComponentResponseDto{
				ComponentCode:      comp.ComponentCode,
				Rate:               comp.Rate,
				LiabilityAccountID: comp.LiabilityAccountID,
			})
		}

		ruleDtos = append(ruleDtos, legal_entity_contracts.LegalEntityTaxCategoryTaxRuleResponseDto{
			LegalEntityTaxRegistrationID: rule.LegalEntityTaxRegistrationID,
			TaxType:                      rule.TaxType,
			Rate:                         rule.Rate,
			MinValue:                     rule.MinValue,
			MaxValue:                     rule.MaxValue,
			EffectiveFrom:                rule.EffectiveFrom,
			EffectiveTo:                  rule.EffectiveTo,
			Components:                   componentDtos,
		})
	}

	return ruleDtos, nil
}
