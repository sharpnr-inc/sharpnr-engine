package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

type PaymentOrderResponseDTO struct {
	ID               uuid.UUID `json:"id"`
	OrderNumber      string    `json:"order_number"`
	LegalEntityID    uuid.UUID `json:"legal_entity_id"`
	PaymentAccountID uuid.UUID `json:"payment_account_id"`

	// Payer Info
	PayerType string  `json:"payer_type"`
	PayerName string  `json:"payer_name"`
	PayerRef  *string `json:"payer_ref,omitempty"`

	// Currency & Totals
	Currency    string  `json:"currency"`
	TotalAmount float64 `json:"total_amount"`
	PaidAmount  float64 `json:"paid_amount"`
	AmountPaise int64   `json:"amount_paise"`
	PaidPaise   int64   `json:"paid_paise"`

	Status      string     `json:"status"`
	Source      string     `json:"source"`
	SourceRefID *string    `json:"source_ref_id,omitempty"`
	SourceMeta  *string    `json:"source_meta,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
