package payment_engine_repo

import (
	"context"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

// -----------------------------------------------
// PRIMARY INTERFACE
// This is the ONLY thing business services import
// Hotel, Restaurant, Bar, Orchestrator all depend
// on this — never on the concrete impl
// -----------------------------------------------
type PaymentEngineRepoInterface interface {
	// Order lifecycle

	GetGatewayConfigByID(ctx context.Context, configID uuid.UUID) (*payment_engine_models.PaymentGatewayConfig, error)
	GetActiveGatewayConfig(ctx context.Context, tenantID uuid.UUID, legalEntityID *uuid.UUID, providerCode string) (*payment_engine_models.PaymentGatewayConfig, error)
}
