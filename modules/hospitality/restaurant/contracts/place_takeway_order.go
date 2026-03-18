package restaurant_contracts

import (
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

// ─── INPUT ───────────────────────────────────────────────────────────────────

type PlaceTakeawayOrderRequest struct {
	OrganizationID uuid.UUID
	RestaurantID   uuid.UUID
	LegalEntityID  uuid.UUID

	// TAKEAWAY, DELIVERY
	Type string

	// Customer
	CustomerID    *uuid.UUID // nil for walk-in
	CustomerName  string
	CustomerPhone string
	CustomerEmail *string

	// Delivery only — nil for TAKEAWAY
	DeliveryAddress *string
	DeliveryNote    *string

	// Drives CGST/SGST vs IGST on invoice
	CustomerStateCode *string

	// Items from cart
	Items []TakeawayOrderItemRequest

	// CASH, CARD, UPI
	// Not PHONEPE/RAZORPAY — those are gateway providers, not payment methods
	// UPI method resolves the correct gateway from PaymentGatewayConfig at runtime
	PaymentMethod string

	// Staff who placed this order
	CreatedBy       string
	CreatedByEmplID uuid.UUID

	// REMOVED: CallbackURL  — static in PaymentGatewayConfig.WebhookURL
	// REMOVED: RedirectURL  — counter app, no browser redirect
	// REMOVED: PayerPhone   — staff app, customer not interacting with system
	// REMOVED: GatewayCode  — resolved automatically from PaymentGatewayConfig
}

type TakeawayOrderItemRequest struct {
	MenuItemID uuid.UUID
	Name       string // snapshot from menu at order time
	Quantity   float64
	Rate       float64 // snapshot from menu at order time

	TaxCategoryID *uuid.UUID
	HSNCode       *string
	Notes         *string

	Variants []TakeawayOrderItemVariantRequest
	AddOns   []TakeawayOrderItemAddOnRequest
}

type TakeawayOrderItemVariantRequest struct {
	VariantID      uuid.UUID
	VariantGroupID uuid.UUID
	Name           string
	PriceAtOrder   float64
}

type TakeawayOrderItemAddOnRequest struct {
	AddOnID      uuid.UUID
	Name         string
	Quantity     int
	PriceAtOrder float64
}

// ─── OUTPUT ──────────────────────────────────────────────────────────────────

type PlaceTakeawayOrderResponse struct {
	OrderID     uuid.UUID
	OrderNumber string
	Type        string

	// CASH    → "CONFIRMED"         order in kitchen, payment done
	// CARD    → "PENDING_PAYMENT"   waiting for staff to confirm terminal
	// UPI     → "PENDING_PAYMENT"   waiting for customer to scan QR
	Status string

	InvoiceID     uuid.UUID
	InvoiceNumber string
	Total         float64

	// Payment status
	// CASH → "PAID"
	// CARD → "PENDING_CARD_CONFIRMATION"
	// UPI  → "PENDING"
	PaymentStatus string

	// UPI only — shown on second screen for customer to scan
	// nil for CASH and CARD
	QRData *string

	// UPI only — when QR expires (typically 10 minutes from PhonePe)
	// Frontend uses this to show countdown timer on second screen
	QRExpiresAt *string
}

// ─── CARD CONFIRMATION ───────────────────────────────────────────────────────
// Staff hits "Card Approved" after POS terminal confirms

type ConfirmCardPaymentRequest struct {
	// For takeaway flow
	OrderID *uuid.UUID

	// For dining session flow
	SessionID *uuid.UUID

	// Reference from POS terminal — printed on terminal receipt
	// e.g. approval code, terminal txn ID
	TerminalRef *string

	CreatedBy string
}

// ─── WEBHOOK CONFIRMED ───────────────────────────────────────────────────────
// Used internally when PhonePe webhook fires
// Built by sero event handler — not exposed to API

type TakeawayPaymentConfirmedRequest struct {
	// Echoed back from PaymentOrder.SourceMeta
	SourceMeta string

	// Confirmed payment record from payment engine
	Payment *payment_engine_models.Payment
}

// ─── STATUS UPDATES ──────────────────────────────────────────────────────────

type UpdateTakeawayStatusRequest struct {
	OrderID uuid.UUID

	// PREPARING → set by kitchen when they start
	// READY     → set by kitchen when done
	// COMPLETED → set by staff when customer picks up / delivered
	// CANCELLED → set by staff
	Status string

	ByRef string
}

type CancelTakeawayOrderRequest struct {
	OrderID uuid.UUID
	Reason  string
	ByRef   string
}

// ─── VALIDATION ──────────────────────────────────────────────────────────────

func (r *PlaceTakeawayOrderRequest) Validate() error {
	if r.OrganizationID == uuid.Nil {
		return fmt.Errorf("organization ID is required")
	}
	if r.RestaurantID == uuid.Nil {
		return fmt.Errorf("restaurant ID is required")
	}
	if r.LegalEntityID == uuid.Nil {
		return fmt.Errorf("legal entity ID is required")
	}
	if r.CreatedByEmplID == uuid.Nil {
		return fmt.Errorf("employee ID is required")
	}
	if r.CustomerName == "" {
		return fmt.Errorf("customer name is required")
	}
	if r.CustomerPhone == "" {
		return fmt.Errorf("customer phone is required")
	}
	if r.Type != "TAKEAWAY" && r.Type != "DELIVERY" {
		return fmt.Errorf("type must be TAKEAWAY or DELIVERY")
	}
	if r.Type == "DELIVERY" && r.DeliveryAddress == nil {
		return fmt.Errorf("delivery address is required for delivery orders")
	}
	if len(r.Items) == 0 {
		return fmt.Errorf("order must have at least one item")
	}
	for i, item := range r.Items {
		if item.MenuItemID == uuid.Nil {
			return fmt.Errorf("item[%d]: menu item ID is required", i)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("item[%d]: quantity must be greater than zero", i)
		}
		if item.Rate <= 0 {
			return fmt.Errorf("item[%d]: rate must be greater than zero", i)
		}
		if item.Name == "" {
			return fmt.Errorf("item[%d]: name is required", i)
		}
	}
	if err := validatePaymentMethod(r.PaymentMethod); err != nil {
		return err
	}
	return nil
}

func validatePaymentMethod(method string) error {
	switch method {
	case "CASH", "CARD", "UPI":
		return nil
	default:
		return fmt.Errorf(
			"payment method must be CASH, CARD or UPI — got '%s'",
			method,
		)
	}
}
