package payment_engine_contracts

import "github.com/google/uuid"

type CreateOrderRequest struct {
	TenantID         uuid.UUID
	PayerType        string  // CUSTOMER, GUEST, WALK_IN
	PayerRef         *string // caller's customer ID — opaque
	PayerName        string
	CurrencyCode     string
	AmountPaise      int64
	Source           string  // INVOICE, POS, ADVANCE
	SourceRefID      *string // invoice ID, session ID — opaque string
	SourceMeta       *string // optional JSON echoed back in webhooks
	ExpiresInMinutes int
	CreatedBy        string
}

type PayRequest struct {
	PaymentOrderID  uuid.UUID
	AmountPaise     int64
	Method          string  // CASH, UPI, CARD, PHONEPE, RAZORPAY
	ReferenceNumber *string // UTR, cheque no — offline only
	PayerPhone      *string // for UPI / gateway
	PayerEmail      *string
	CallbackURL     string  // where gateway posts webhook
	RedirectURL     *string // where user lands after online payment
	Notes           *string
	CreatedBy       string
}

type ConfirmPaymentRequest struct {
	GatewayOrderRef string
	GatewayTxnID    string
	RawCallbackBody string
	Headers         map[string]string
}

type AllocateRequest struct {
	PaymentID      uuid.UUID
	TargetType     string // INVOICE, ADVANCE, CREDIT_NOTE
	TargetID       string // invoice UUID as string — opaque
	AllocatedPaise int64
	CreatedBy      string
}

type RefundRequest struct {
	PaymentID   uuid.UUID
	AmountPaise int64
	Reason      string
	InitiatedBy string
	SourceRef   *string // cancellation ID, complaint ID — opaque
}

type AddGatewayRequest struct {
	TenantID     uuid.UUID
	ProviderCode string
	DisplayName  string
	Config       map[string]string // will be encrypted before save
	IsTestMode   bool
}

type UpdateGatewayRequest struct {
	DisplayName *string
	Config      map[string]string
	IsTestMode  *bool
}
