// Package payment_gateway defines the core abstraction layer between the
// payment engine and any external payment provider (PhonePe, Razorpay, Cash etc.)
//
// ARCHITECTURE:
// The payment engine never talks to a gateway directly.
// It always goes through this interface.
// This means adding a new gateway = implement this interface only.
// Nothing else in the engine changes.
//
// FLOW:
//   PaymentService
//       │
//       ▼
//   GatewayResolver  →  resolves correct provider from registry
//       │
//       ▼
//   GatewayProvider  →  CreateOrder / VerifyPayment / InitiateRefund / FetchStatus
//       │
//       ▼
//   External Gateway API (PhonePe, Razorpay) OR self-confirm (Cash, Card)
//
// COUNTER-SPECIFIC DESIGN DECISIONS:
//   - No RedirectURL / CallbackURL per transaction
//     These are registered ONCE in PaymentGatewayConfig.WebhookURL during setup
//     and injected into Config map by the resolver before each call
//   - GatewayOrderResponse returns QRData (not PaymentURL)
//     Because this is a counter/staff app — customer never opens a browser
//     QR is shown on second screen, customer scans with their UPI app
//   - CARD payments bypass this interface entirely
//     POS terminal is external hardware — staff confirms manually
//     PaymentService handles CARD as an offline method

package payment_gateway

import (
	"context"
	"time"
	
)

// ─────────────────────────────────────────────────────────────────────────────
// CORE INTERFACE
// Every gateway must implement all four methods.
// The engine only ever sees this interface — never a concrete provider.
// ─────────────────────────────────────────────────────────────────────────────

type GatewayProvider interface {
	// Code returns the unique provider identifier.
	// Must match the ProviderCode stored in PaymentGatewayConfig.
	// Examples: "CASH", "PHONEPE", "RAZORPAY"
	Code() string

	// CreateOrder initiates a payment request on the gateway.
	// For UPI (PhonePe/Razorpay): calls gateway API, returns QR string
	// For Cash: self-confirms, echoes back internal order ID
	// For Card: not called — handled by POS terminal externally
	CreateOrder(ctx context.Context, req GatewayOrderRequest) (*GatewayOrderResponse, error)

	// VerifyPayment validates and processes an inbound webhook callback.
	// Called from webhook_handlers.go when gateway POSTs payment confirmation.
	// Verifies HMAC signature, decodes payload, maps to internal status.
	VerifyPayment(ctx context.Context, req GatewayVerifyRequest) (*GatewayVerifyResponse, error)

	// InitiateRefund requests a refund on the gateway side.
	// For Cash: self-confirms (manual physical refund)
	// For UPI: calls gateway refund API
	InitiateRefund(ctx context.Context, req GatewayRefundRequest) (*GatewayRefundResponse, error)

	// FetchStatus polls the gateway for current transaction status.
	// Used for reconciliation and missed webhook recovery.
	// Note: requires Config injection — call via service layer not directly.
	FetchStatus(ctx context.Context, gatewayTxnID string, config map[string]string) (*GatewayStatusResponse, error)
}

// ─────────────────────────────────────────────────────────────────────────────
// REQUEST TYPES
// Shared across all providers.
// Provider-specific fields are handled inside each provider package — never here.
// ─────────────────────────────────────────────────────────────────────────────

// GatewayOrderRequest is the input to CreateOrder.
// Config carries decrypted provider credentials + webhook URL.
// The resolver builds this from PaymentGatewayConfig before calling the provider.
type GatewayOrderRequest struct {
	// InternalOrderID is your Payment.ID — used as MerchantTransactionID on gateway.
	// This is how you map a gateway webhook back to your internal payment record.
	InternalOrderID string

	// AmountPaise is always in the smallest currency unit (paise for INR).
	// Never use float for money — avoids rounding bugs.
	AmountPaise  int64
	CurrencyCode string // "INR"

	// Config carries decrypted key-value pairs from PaymentGatewayConfig.ConfigEncrypted
	// PLUS the WebhookURL injected by the resolver.
	// PhonePe:  { "merchant_id", "salt_key", "salt_index", "webhook_url", "is_test_mode" }
	// Razorpay: { "key_id", "key_secret", "webhook_url", "is_test_mode" }
	Config map[string]string

	// Meta is optional caller context echoed back in webhooks via SourceMeta.
	// Used by restaurant service to identify session/order on webhook confirmation.
	// Example: { "session_id": "...", "invoice_id": "...", "flow": "DINING_SESSION" }
	Meta map[string]string
}

// GatewayRefundRequest is the input to InitiateRefund.
type GatewayRefundRequest struct {
	// GatewayTxnID is the gateway's own transaction ID from the original payment.
	// Set from Payment.GatewayTxnRef.
	GatewayTxnID string

	// InternalRefundID is your Refund.ID — used as the refund reference on gateway.
	InternalRefundID string

	// AmountPaise for this refund — can be partial.
	RefundAmountPaise int64

	// TotalAmountPaise is the original payment amount.
	// Required by some gateways (e.g. PhonePe) for validation.
	TotalAmountPaise int64

	Reason string
	Config map[string]string
}

// GatewayVerifyRequest is the input to VerifyPayment.
// Built from the raw inbound webhook in webhook_handlers.go.
type GatewayVerifyRequest struct {
	// GatewayTxnID and GatewayOrderID come from the webhook payload.
	GatewayTxnID   string
	GatewayOrderID string // your InternalOrderID echoed back by gateway

	AmountPaise int64

	// RawCallbackBody is the exact raw bytes from the gateway POST.
	// Used for HMAC signature verification — must not be modified.
	RawCallbackBody string

	// Headers from the gateway webhook request.
	// PhonePe sends X-VERIFY, Razorpay sends X-Razorpay-Signature.
	Headers map[string]string

	Config map[string]string
}

// ─────────────────────────────────────────────────────────────────────────────
// RESPONSE TYPES
// ─────────────────────────────────────────────────────────────────────────────

// GatewayOrderResponse is the output of CreateOrder.
type GatewayOrderResponse struct {
	// ProviderOrderID is the gateway's own reference for this transaction.
	// Stored in Payment.GatewayOrderRef.
	ProviderOrderID string

	// QRData is the UPI QR string for counter payments.
	// Shown on the second screen — customer scans with any UPI app.
	// nil for Cash and Card (no QR needed).
	// Format: "upi://pay?pa=merchant@upi&pn=MerchantName&am=450.00&tr=TXN123&tn=OrderDesc"
	QRData *string

	// ExpiresAt is when the QR/transaction expires on the gateway side.
	// Used to show a countdown timer on the second screen.
	ExpiresAt *time.Time

	// RawResponse is the full gateway JSON response — stored as snapshot in Payment.GatewayResponse.
	// Never updated after initial save — immutable audit record.
	RawResponse string
}

// GatewayVerifyResponse is the output of VerifyPayment.
type GatewayVerifyResponse struct {
	// Status maps to your internal PaymentStatus constants.
	// "SUCCESS", "FAILED", "PROCESSING"
	Status string

	GatewayTxnID   string
	GatewayOrderID string

	// PaidAmountPaise is what the gateway actually received.
	// Should match AmountPaise — discrepancy = fraud check needed.
	PaidAmountPaise int64

	// PaidAt is the actual settlement timestamp from the gateway.
	PaidAt time.Time

	RawResponse string
}

// GatewayRefundResponse is the output of InitiateRefund.
type GatewayRefundResponse struct {
	// GatewayRefundID is the gateway's reference for this refund.
	// Stored in Refund.GatewayRefundRef.
	GatewayRefundID string

	// Status: "SUCCESS", "PROCESSING", "FAILED"
	// Cash refunds are always "SUCCESS" (manual handback).
	Status string

	RawResponse string
}

// GatewayStatusResponse is the output of FetchStatus.
type GatewayStatusResponse struct {
	Status          string
	GatewayTxnID    string
	GatewayOrderID  string
	PaidAmountPaise int64
	RawResponse     string
}
