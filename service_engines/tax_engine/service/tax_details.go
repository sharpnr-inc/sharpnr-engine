package tax_engine_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// func (s *taxEngineService) GetTaxDetails(
//     ctx context.Context,
//     taxCategoryID uuid.UUID,
//     legalEntityID uuid.UUID,
// ) ([]tax_engine_contracts.TaxComponentBreakdown, error) {

//     var res []tax_engine_contracts.TaxComponentBreakdown

//     err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
//         txCtx := dbtx.InjectTx(ctx, tx)

//         // 1. Fetch tax category details from the Legal Entity Service
//         taxDetails, err := s.legalEntityService.GetTaxCategoryDetails(txCtx, taxCategoryID, legalEntityID)
//         if err != nil {
//             return err
//         }

//         // 2. Identify the active rule based on the current time
//         now := time.Now()
//         var activeRule *legal_entity_contracts.LegalEntityTaxCategoryTaxRuleResponseDto

//         for _, rule := range taxDetails.Rules {
//             if now.After(rule.EffectiveFrom) && (rule.EffectiveTo == nil || now.Before(*rule.EffectiveTo)) {
//                 activeRule = &rule
//                 break
//             }
//         }

//         if activeRule == nil {
//             return errors.New("no active tax rule found for the current date")
//         }

//         // 3. Map components from the active rule to the response breakdown
//         for _, comp := range activeRule.Components {
//             res = append(res, tax_engine_contracts.TaxComponentBreakdown{
//                 Name: comp.ComponentCode,
//                 Rate: comp.Rate,
//             })
//         }

//         return nil
//     })

//     if err != nil {
//         return nil, err
//     }

//     return res, nil
// }

func (s *taxEngineService) GetTaxDetails(
	ctx context.Context,
	taxCategoryID uuid.UUID,
	legalEntityID uuid.UUID,
) ([]tax_engine_contracts.TaxComponentBreakdown, error) {

	var res []tax_engine_contracts.TaxComponentBreakdown

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch tax category details from the Legal Entity Service
		taxDetails, err := s.legalEntityService.GetTaxCategoryDetails(txCtx, taxCategoryID, legalEntityID)
		if err != nil {
			return err
		}

		// // 2. Simple Check: Ensure there is at least one rule available
		// if len(taxDetails.Rules) == 0 {
		// 	return errors.New("no tax rules found for this category")
		// }

		// 3. Take components from the first rule (ignoring effective dates for now)
		firstRule := taxDetails.Rules[0]

		for _, comp := range firstRule.Components {
			res = append(res, tax_engine_contracts.TaxComponentBreakdown{
				Name: comp.ComponentCode,
				Rate: comp.Rate,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
