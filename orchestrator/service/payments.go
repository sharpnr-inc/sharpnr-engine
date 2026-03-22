package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *orchestratorService) AddPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req payment_engine_contracts.CreatePaymentAccountRequest) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		_, err := c.paymentsEngine.CreatePaymentAccount(txCtx, orgID, emplID, req)
		if err != nil {
			return err
		}

		return nil
	})
}

func (c *orchestratorService) GetAllPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]payment_engine_contracts.PaymentAccountResponseDto, error) {
	var results []payment_engine_contracts.PaymentAccountResponseDto

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch the accounts from the payment engine
		accounts, err := c.paymentsEngine.ListPaymentAccounts(txCtx, orgID)
		if err != nil {
			return err
		}

		// 2. Map the domain models to the DTOs
		for _, acc := range accounts {
			dto := payment_engine_contracts.PaymentAccountResponseDto{
				Id:            acc.ID,
				LegalEntityID: acc.LegalEntityID,
				Name:          acc.AccountRef, // Mapping AccountRef to Name
				Currency:      acc.CurrencyCode,
				IsActive:      acc.IsActive,
			}
			results = append(results, dto)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
