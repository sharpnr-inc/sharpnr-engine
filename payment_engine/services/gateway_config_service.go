package payment_engine_service

import (
	"context"
	"fmt"

	contracts "sharpnr-engine/payment_engine/contracts"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

func (s *paymentEngine) AddGatewayConfig(
	ctx context.Context,
	req contracts.AddGatewayConfigRequest,
) (*models.PaymentGatewayConfig, error) {

	account, err := s.paymentRepo.GetPaymentAccountByID(ctx, req.PaymentAccountID)
	if err != nil {
		return nil, fmt.Errorf("payment account not found: %w", err)
	}

	// if !s.gatewayRegistry.IsRegistered(req.ProviderCode) {
	// 	return nil, fmt.Errorf("provider '%s' is not registered", req.ProviderCode)
	// }

	existing, err := s.paymentRepo.GetActiveGatewayConfig(ctx, req.PaymentAccountID, req.ProviderCode)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("gateway '%s' is already configured for this account", req.ProviderCode)
	}

	encrypted, err := encryptConfig(req.Config)
	if err != nil {
		return nil, fmt.Errorf("encrypt gateway config: %w", err)
	}

	config := &models.PaymentGatewayConfig{
		ID:               uuid.New(),
		TenantID:         account.TenantID,
		LegalEntityID:    account.LegalEntityID,
		PaymentAccountID: req.PaymentAccountID,
		ProviderCode:     req.ProviderCode,
		DisplayName:      req.DisplayName,
		ConfigEncrypted:  encrypted,
		WebhookURL:       req.WebhookURL,
		IsTestMode:       req.IsTestMode,
		IsActive:         true,
	}

	if err := s.paymentRepo.CreateGatewayConfig(ctx, config); err != nil {
		return nil, fmt.Errorf("create gateway config: %w", err)
	}

	config.ConfigEncrypted = "" // never return credentials
	return config, nil
}

func (s *paymentEngine) ListGatewayConfigs(
	ctx context.Context,
	paymentAccountID uuid.UUID,
) ([]*models.PaymentGatewayConfig, error) {

	configs, err := s.paymentRepo.ListGatewayConfigsByAccount(ctx, paymentAccountID)
	if err != nil {
		return nil, fmt.Errorf("list gateway configs: %w", err)
	}
	for _, c := range configs {
		c.ConfigEncrypted = "" // never leak credentials
	}
	return configs, nil
}

func (s *paymentEngine) UpdateGatewayConfig(
	ctx context.Context,
	id uuid.UUID,
	req contracts.UpdateGatewayConfigRequest,
) error {

	existing, err := s.paymentRepo.GetGatewayConfigByID(ctx, id)
	if err != nil {
		return fmt.Errorf("gateway config not found: %w", err)
	}

	encrypted, err := encryptConfig(req.Config)
	if err != nil {
		return fmt.Errorf("encrypt updated config: %w", err)
	}

	webhookURL := existing.WebhookURL
	if req.WebhookURL != "" {
		webhookURL = req.WebhookURL
	}

	return s.paymentRepo.UpdateGatewayConfig(ctx, id, encrypted, webhookURL)
}

func (s *paymentEngine) ToggleGatewayConfig(
	ctx context.Context,
	id uuid.UUID,
	active bool,
) error {

	if _, err := s.paymentRepo.GetGatewayConfigByID(ctx, id); err != nil {
		return fmt.Errorf("gateway config not found: %w", err)
	}
	return s.paymentRepo.ToggleGatewayConfig(ctx, id, active)
}

func (s *paymentEngine) RemoveGatewayConfig(
	ctx context.Context,
	id uuid.UUID,
) error {

	if _, err := s.paymentRepo.GetGatewayConfigByID(ctx, id); err != nil {
		return fmt.Errorf("gateway config not found: %w", err)
	}
	return s.paymentRepo.DeleteGatewayConfig(ctx, id)
}
