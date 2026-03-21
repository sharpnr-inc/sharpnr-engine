// payment_engine/contracts/input_contracts.go

package payment_engine_contracts

import "github.com/google/uuid"

// ─── Account ──────────────────────────────────────────────────────────────────

type CreatePaymentAccountRequest struct {
	TenantID      uuid.UUID `json:"tenant_id" binding:"required"`
	LegalEntityID uuid.UUID `json:"legal_entity_id" binding:"required"`
	AccountRef    string    `json:"account_ref" binding:"required"`
	CurrencyCode  string    `json:"currency_code"`
}

// ─── Gateway Config ───────────────────────────────────────────────────────────

type AddGatewayConfigRequest struct {
	PaymentAccountID uuid.UUID         `json:"payment_account_id" binding:"required"`
	ProviderCode     string            `json:"provider_code" binding:"required"`
	DisplayName      string            `json:"display_name" binding:"required"`
	Config           map[string]string `json:"config" binding:"required"` // plain — service encrypts
	WebhookURL       string            `json:"webhook_url" binding:"required"`
	IsTestMode       bool              `json:"is_test_mode"`
}

type UpdateGatewayConfigRequest struct {
	Config     map[string]string `json:"config"`
	WebhookURL string            `json:"webhook_url"`
}

// ─── Payment Order ────────────────────────────────────────────────────────────

type CreateOrderRequest struct {
	TenantID         uuid.UUID `json:"tenant_id" binding:"required"`
	LegalEntityID    uuid.UUID `json:"legal_entity_id" binding:"required"`
	PaymentAccountID uuid.UUID `json:"payment_account_id" binding:"required"`
	PayerType        string    `json:"payer_type" binding:"required"`
	PayerRef         *string   `json:"payer_ref"`
	PayerName        string    `json:"payer_name" binding:"required"`
	CurrencyCode     string    `json:"currency_code"`
	AmountPaise      int64     `json:"amount_paise" binding:"required"`
	Source           string    `json:"source"`
	SourceRefID      *string   `json:"source_ref_id"`
	SourceMeta       *string   `json:"source_meta"`
	ExpiresInMinutes int       `json:"expires_in_minutes"`
	CreatedBy        string    `json:"created_by" binding:"required"`
}

// ─── Payment ──────────────────────────────────────────────────────────────────

type PayRequest struct {
	PaymentOrderID  uuid.UUID `json:"payment_order_id" binding:"required"`
	AmountPaise     int64     `json:"amount_paise" binding:"required"`
	Method          string    `json:"method" binding:"required"`
	ReferenceNumber *string   `json:"reference_number"` // UTR, cheque no — offline only
	PayerPhone      *string   `json:"payer_phone"`
	PayerEmail      *string   `json:"payer_email"`
	CallbackURL     string    `json:"callback_url"`
	RedirectURL     *string   `json:"redirect_url"`
	Notes           *string   `json:"notes"`
	CreatedBy       string    `json:"created_by" binding:"required"`
}

type ConfirmPaymentRequest struct {
	GatewayOrderRef string            `json:"gateway_order_ref"`
	GatewayTxnID    string            `json:"gateway_txn_id"`
	RawCallbackBody string            `json:"raw_callback_body"`
	Headers         map[string]string `json:"headers"`
}

// ─── Allocation ───────────────────────────────────────────────────────────────

type AllocateRequest struct {
	PaymentID      uuid.UUID `json:"payment_id" binding:"required"`
	TargetType     string    `json:"target_type" binding:"required"`
	TargetID       string    `json:"target_id" binding:"required"`
	AllocatedPaise int64     `json:"allocated_paise" binding:"required"`
	CreatedBy      string    `json:"created_by" binding:"required"`
}

type ReverseAllocationRequest struct {
	AllocationID uuid.UUID `json:"allocation_id" binding:"required"`
	By           string    `json:"by" binding:"required"`
	Note         string    `json:"note"`
}

// ─── Refund ───────────────────────────────────────────────────────────────────

type RefundRequest struct {
	PaymentID   uuid.UUID `json:"payment_id" binding:"required"`
	AmountPaise int64     `json:"amount_paise" binding:"required"`
	Reason      string    `json:"reason" binding:"required"`
	InitiatedBy string    `json:"initiated_by" binding:"required"`
	SourceRef   *string   `json:"source_ref"`
}

// ─── Webhook Endpoint ─────────────────────────────────────────────────────────

type RegisterWebhookRequest struct {
	TenantID         uuid.UUID `json:"tenant_id" binding:"required"`
	URL              string    `json:"url" binding:"required"`
	SubscribedEvents []string  `json:"subscribed_events"` // empty = all
}
