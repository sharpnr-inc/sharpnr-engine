package restaurant_events

import (
	"errors"
	"sharpnr-engine/sero"
	"time"

	"github.com/google/uuid"
)

const (
	OrderCreatedInRestaurant sero.SeroEngineEventType = "ORDER_CREATED_IN_RESTAURANT"
	OrderStatusUpdated       sero.SeroEngineEventType = "ORDER_STATUS_UPDATED"
	RestaurantPaymentSettled sero.SeroEngineEventType = "RESTAURANT_PAYMENT_SETTLED"
)

type OrderStatusUpdatePayload struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
}

type OrderCreatedPayload struct {
	OrderID            uuid.UUID           `json:"order_id"`
	OrganizationID     uuid.UUID           `json:"organization_id"`
	RestaurantID       uuid.UUID           `json:"restaurant_id"`
	PreferredKitchenID *uuid.UUID          `json:"preferred_kitchen_id,omitempty"`
	SessionID          uuid.UUID           `json:"session_id"`
	Items              []OrderItemEventDTO `json:"items"`
}

type OrderItemEventDTO struct {
	MenuItemID   uuid.UUID         `json:"menu_item_id"`
	MenuItemName string            `json:"menu_item_name"`
	Quantity     int               `json:"quantity"`
	Notes        string            `json:"notes"`
	Variants     []VariantEventDTO `json:"variants"`
	Addons       []AddonEventDTO   `json:"addons"`
}

type VariantEventDTO struct {
	VariantID    uuid.UUID `json:"variant_id"`
	VariantName  string    `json:"variant_name"`
	PriceAtOrder float64   `json:"price_at_order"`
}

type AddonEventDTO struct {
	AddOnID      uuid.UUID `json:"addon_id"`
	Name         string    `json:"name"`
	Qty          int       `json:"qty"`
	PriceAtOrder float64   `json:"price_at_order"`
}

/* ---------------------------------- <***> --------------------------------- */

type RestaurantPaymentSettledDto struct {
	InvoiceID      uuid.UUID
	PaymentID      uuid.UUID
	OrganizationID uuid.UUID
	LegalEntityID  uuid.UUID
	AmountPaise    float64
	PaidAt         time.Time
}

func (rps *RestaurantPaymentSettledDto) ValidateDto() error {

	if rps.InvoiceID == uuid.Nil {
		return errors.New("invoice_id is required")
	}

	if rps.PaymentID == uuid.Nil {
		return errors.New("payment_id is required")
	}

	if rps.OrganizationID == uuid.Nil {
		return errors.New("organization_id is required")
	}

	if rps.LegalEntityID == uuid.Nil {
		return errors.New("legal_entity_id is required")
	}

	if rps.AmountPaise <= 0 {
		return errors.New("amount_paise must be greater than zero")
	}

	if rps.PaidAt.IsZero() {
		return errors.New("paid_at is required")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
