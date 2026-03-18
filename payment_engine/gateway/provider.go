package payment_gateway

import (
	"context"
	"time"
)

// -----------------------------------------------
// CORE INTERFACE
// Every gateway must implement this — nothing else
// in the engine cares which gateway is being used
// -----------------------------------------------
type GatewayProvider interface {
	Code() string // "CASH", "PHONEPE", "RAZORPAY"

	CreateOrder(ctx context.Context, req GatewayOrderRequest) (*GatewayOrderResponse, error)
	VerifyPayment(ctx context.Context, req GatewayVerifyRequest) (*GatewayVerifyResponse, error)
	InitiateRefund(ctx context.Context, req GatewayRefundRequest) (*GatewayRefundResponse, error)
	FetchStatus(ctx context.Context, gatewayTxnID string) (*GatewayStatusResponse, error)
}

// -----------------------------------------------
// SHARED REQUEST / RESPONSE TYPES
// These are the only types the engine passes into
// a provider — no provider-specific types leak out
// -----------------------------------------------

type GatewayOrderRequest struct {
	InternalOrderID string // your PaymentOrder.ID
	AmountPaise     int64  // always smallest unit
	CurrencyCode    string // "INR"
	PayerName       string
	PayerPhone      *string
	PayerEmail      *string
	CallbackURL     string            // where gateway sends webhook
	RedirectURL     *string           // where user lands after payment (web flow)
	Config          map[string]string // decrypted provider config values
	Meta            map[string]string // optional caller metadata echoed back
}

type GatewayOrderResponse struct {
	ProviderOrderID string  // gateway's own order/session ref
	PaymentURL      *string // redirect URL for web payments
	QRCode          *string // UPI QR string
	DeepLink        *string // UPI deep link for mobile
	ExpiresAt       *time.Time
	RawResponse     string // full gateway JSON — snapshotted in payments.gateway_response
}

type GatewayVerifyRequest struct {
	GatewayTxnID    string
	GatewayOrderID  string
	AmountPaise     int64
	RawCallbackBody string            // raw webhook body for signature verification
	Headers         map[string]string // webhook headers — needed for HMAC checks
	Config          map[string]string
}

type GatewayVerifyResponse struct {
	Status          string // maps to your PaymentStatus constants: SUCCESS, FAILED, PENDING
	GatewayTxnID    string
	GatewayOrderID  string
	PaidAmountPaise int64
	PaidAt          time.Time
	RawResponse     string
}

type GatewayRefundRequest struct {
	GatewayTxnID      string
	InternalRefundID  string // your Refund.ID — some gateways need a unique refund ref
	RefundAmountPaise int64
	TotalAmountPaise  int64 // original payment amount — required by some gateways
	Reason            string
	Config            map[string]string
}

type GatewayRefundResponse struct {
	GatewayRefundID string // gateway's refund transaction ID
	Status          string // SUCCESS, PENDING, FAILED
	RawResponse     string
}

type GatewayStatusResponse struct {
	Status          string
	GatewayTxnID    string
	GatewayOrderID  string
	PaidAmountPaise int64
	RawResponse     string
}
