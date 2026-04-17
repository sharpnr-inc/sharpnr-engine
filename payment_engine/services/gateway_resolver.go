package payment_engine_service

import (
	"context"
	"fmt"

	payment_gateway "sharpnr-engine/payment_engine/gateway"

	"github.com/google/uuid"
)

type resolvedGateway struct {
	Provider payment_gateway.GatewayProvider
	Config   map[string]string
	ConfigID uuid.UUID
}

func (s *paymentEngine) resolveGateway(
	ctx context.Context,
	paymentAccountID uuid.UUID,
	providerCode string,
) (*resolvedGateway, error) {

	config, err := s.paymentRepo.GetActiveGatewayConfig(ctx, paymentAccountID, providerCode)
	if err != nil {
		return nil, fmt.Errorf("no active gateway config for provider '%s': %w", providerCode, err)
	}

	decrypted, err := decryptConfig(config.ConfigEncrypted)
	if err != nil {
		return nil, fmt.Errorf("decrypt gateway config: %w", err)
	}

	provider, err := s.gatewayRegistry.Resolve(providerCode)
	if err != nil {
		return nil, err
	}

	return &resolvedGateway{
		Provider: provider,
		Config:   decrypted,
		ConfigID: config.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) resolveGatewayByConfigID(
	ctx context.Context,
	configID uuid.UUID,
) (*resolvedGateway, error) {

	config, err := s.paymentRepo.GetGatewayConfigByID(ctx, configID)
	if err != nil {
		return nil, fmt.Errorf("gateway config not found: %w", err)
	}
	if !config.IsActive {
		return nil, fmt.Errorf("gateway config '%s' is inactive", config.ProviderCode)
	}

	decrypted, err := decryptConfig(config.ConfigEncrypted)
	if err != nil {
		return nil, fmt.Errorf("decrypt gateway config: %w", err)
	}

	provider, err := s.gatewayRegistry.Resolve(config.ProviderCode)
	if err != nil {
		return nil, err
	}

	return &resolvedGateway{
		Provider: provider,
		Config:   decrypted,
		ConfigID: config.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */
