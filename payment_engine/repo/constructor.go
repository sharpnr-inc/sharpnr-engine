package payment_engine_repo

import (
	"context"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

// GetActiveGatewayConfig implements [PaymentEngineRepoInterface].
func (p *paymentRepo) GetActiveGatewayConfig(ctx context.Context, tenantID uuid.UUID, legalEntityID *uuid.UUID, providerCode string) (*payment_engine_models.PaymentGatewayConfig, error) {
	panic("unimplemented")
}

// GetGatewayConfigByID implements [PaymentEngineRepoInterface].
func (p *paymentRepo) GetGatewayConfigByID(ctx context.Context, configID uuid.UUID) (*payment_engine_models.PaymentGatewayConfig, error) {
	panic("unimplemented")
}

// payment Constructor
func NewPaymentRepository(db *gorm.DB) *paymentRepo {
	return &paymentRepo{
		db: db,
	}
}
