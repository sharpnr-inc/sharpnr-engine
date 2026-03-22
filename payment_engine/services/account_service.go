package payment_engine_service

import (
	"context"
	"fmt"

	contracts "sharpnr-engine/payment_engine/contracts"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

func (s *paymentEngine) CreatePaymentAccount(
	ctx context.Context,
	req contracts.CreatePaymentAccountRequest,
) (*models.PaymentAccount, error) {

	existing, err := s.paymentRepo.GetPaymentAccountByLegalEntity(ctx, req.TenantID, req.LegalEntityID)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("a payment account already exists for this legal entity")
	}

	if req.CurrencyCode == "" {
		req.CurrencyCode = "INR"
	}

	account := &models.PaymentAccount{
		ID:            uuid.New(),
		TenantID:      req.TenantID,
		LegalEntityID: req.LegalEntityID,
		AccountRef:    req.AccountRef,
		CurrencyCode:  req.CurrencyCode,
		IsActive:      true,
	}

	if err := s.paymentRepo.CreatePaymentAccount(ctx, account); err != nil {
		return nil, fmt.Errorf("create payment account: %w", err)
	}

	return account, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) ListPaymentAccounts(
	ctx context.Context,
	tenantID uuid.UUID,
) ([]*models.PaymentAccount, error) {

	accounts, err := s.paymentRepo.ListPaymentAccounts(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("list payment accounts: %w", err)
	}
	return accounts, nil
}

func (s *paymentEngine) TogglePaymentAccount(
	ctx context.Context,
	id uuid.UUID,
	active bool,
) error {

	if _, err := s.paymentRepo.GetPaymentAccountByID(ctx, id); err != nil {
		return fmt.Errorf("payment account not found: %w", err)
	}
	return s.paymentRepo.UpdatePaymentAccountStatus(ctx, id, active)
}

/* ---------------------------------- <***> --------------------------------- */
