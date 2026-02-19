package restaurant_events

import (
	"sharpnr-engine/sero"

	"github.com/google/uuid"
)

const (
	OrderCreatedInRestaurant sero.SeroEngineEventType = "ORDER_CREATED_IN_RESTAURANT"
)

type OrderCreatedPayload struct {
	OrderID        uuid.UUID           `json:"order_id"`
	OrganizationID uuid.UUID           `json:"organization_id"`
	RestaurantID   uuid.UUID           `json:"restaurant_id"`
	SessionID      uuid.UUID           `json:"session_id"`
	Items          []OrderItemEventDTO `json:"items"`
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
