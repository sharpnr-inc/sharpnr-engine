package restaurant_contracts

import (
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

// ─── INPUT ───────────────────────────────────────────────────────────────────

type CloseSessionRequest struct {
	SessionID uuid.UUID

	// CASH, CARD, UPI
	PaymentMethod string

	// Walk-in customer override
	// If session has a RestaurantDiningSessionCustomer row already — ignored
	// If no customer row yet (pure walk-in) — used to build snapshot
	CustomerName      string
	CustomerGST       *string
	CustomerStateCode *string // drives CGST/SGST vs IGST

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

// ─── VALIDATION ──────────────────────────────────────────────────────────────

func (r *CloseSessionRequest) Validate() error {
	if r.SessionID == uuid.Nil {
		return fmt.Errorf("session ID is required")
	}

	if err := validatePaymentMethod(r.PaymentMethod); err != nil {
		return err
	}
	return nil
}
