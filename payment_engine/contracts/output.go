package payment_engine_contracts

import (
	"time"

	"github.com/google/uuid"
)

type PaymentAccountResponseDto struct {
	Id            uuid.UUID
	LegalEntityID uuid.UUID
	Name          string
	Currency      string
	IsActive      bool
}

type GatewayConfigResponse struct {
	ID               uuid.UUID `json:"id"`
	PaymentAccountID uuid.UUID `json:"payment_account_id" binding:"required"`
	ProviderCode     string    `json:"provider_code" binding:"required"`
	DisplayName      string    `json:"display_name" binding:"required"`
	IsTestMode       bool      `json:"is_test_mode"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
}
