package legal_entity_service

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) GetTaxCategoryDetails(
	ctx context.Context,
	taxCategoryID uuid.UUID,
	legalEntityID uuid.UUID,
) (legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error) {

	var response legal_entity_contracts.LegalEntityTaxCategoryResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Tax Category
		taxCategory, err := s.legalEntityRepo.GetTaxCategoryByID(txCtx, taxCategoryID)
		if err != nil {
			return err
		}

		// Safety check
		if taxCategory.LegalEntityID != legalEntityID {
			return fmt.Errorf("tax category does not belong to this legal entity")
		}

		// Account Mapping
		accountMapping, err := s.legalEntityRepo.GetTaxCategoryAccountMappingByTaxCategoryID(txCtx, taxCategory.ID)
		if err != nil {
			return err
		}

		// Rules
		rules, err := s.legalEntityRepo.GetActiveRulesByCategory(txCtx, taxCategory.ID)
		if err != nil {
			return err
		}

		ruleDtos, err := s.buildRuleDtos(txCtx, rules)
		if err != nil {
			return err
		}

		// Build response
		response = legal_entity_contracts.LegalEntityTaxCategoryResponseDto{
			ID:           taxCategory.ID,
			CategoryName: taxCategory.Name,
			Code:         taxCategory.Code,
			Description:  taxCategory.Description,
			Rules:        ruleDtos,
		}

		if accountMapping != nil {
			response.AccountingBookID = accountMapping.AccountingBookID
			response.ReceivableAccountID = accountMapping.ReceivableAccountID
			response.IncomeAccountID = accountMapping.IncomeAccountID
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}
