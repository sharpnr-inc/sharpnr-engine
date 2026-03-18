package payment_engine_repo

import (
	"context"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

type PaymentEngineRepoInterface interface {
	GetGatewayConfigByID(ctx context.Context, configID uuid.UUID) (*payment_engine_models.PaymentGatewayConfig, error)

	GetActiveGatewayConfig(ctx context.Context, tenantID uuid.UUID, legalEntityID *uuid.UUID, providerCode string) (*payment_engine_models.PaymentGatewayConfig, error)
}
