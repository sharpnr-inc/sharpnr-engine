package restaurant_contracts

import (
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

// ─── INPUT ───────────────────────────────────────────────────────────────────

type CloseSessionRequest struct {
	// CASH, CARD, UPI
	PaymentMethod    string    `json:"payment_method"`
	PaymentAccountID uuid.UUID `json:"payment_account_id"`

	// Walk-in customer override
	// If session has a RestaurantDiningSessionCustomer row already — ignored
	// If no customer row yet (pure walk-in) — used to build snapshot
	CustomerName  string  `json:"customer_name"`
	CustomerEmail *string `json:"customer_email,omitempty"`
	CustomerPhone *string `json:"customer_phone,omitempty"`
	CustomerGST   *string `json:"customer_gst,omitempty"`

	// if customer exist
	CustomerID *uuid.UUID `json:"customer_id,omitempty"`
}

// ─── OUTPUT ──────────────────────────────────────────────────────────────────

type CloseSessionResponse struct {
	SessionID uuid.UUID

	InvoiceID     uuid.UUID
	InvoiceNumber string
	Total         float64

	// PAID                     → CASH or CARD confirmed
	// PENDING_CARD_CONFIRMATION → waiting for staff to hit "Card Approved"
	// PENDING                  → UPI QR shown, waiting for customer to scan
	PaymentStatus string

	// CLOSED           → CASH or CARD confirmed
	// AWAITING_PAYMENT → CARD or UPI not yet confirmed
	Status string

	// UPI only — pushed to second screen via SSE
	// Also returned here so main screen can show it too
	QRData      *string
	QRExpiresAt *string
}

// ─── CARD CONFIRMATION ───────────────────────────────────────────────────────

type ConfirmSessionCardPaymentRequest struct {
	SessionID uuid.UUID

	// Approval code or reference from POS terminal receipt
	// Optional — some terminals don't provide this
	TerminalRef *string

	CreatedBy       string
	CreatedByEmplID uuid.UUID
}

type ConfirmSessionCardPaymentResponse struct {
	SessionID     uuid.UUID
	InvoiceID     uuid.UUID
	InvoiceNumber string
	Total         float64
	PaymentStatus string // PAID
	Status        string // CLOSED
}

// ─── WEBHOOK CONFIRMED ───────────────────────────────────────────────────────
// Built internally by sero event handler
// NOT exposed to API

type SessionPaymentConfirmedRequest struct {
	SourceMeta string // JSON from PaymentOrder.SourceMeta
	Payment    *payment_engine_models.Payment
}
