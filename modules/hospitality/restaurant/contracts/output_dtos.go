package restaurant_contracts

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	"time"

	"github.com/google/uuid"
)

type DiningSessionOrders_resp_dto struct {
	DiningSessionID uuid.UUID        `json:"dining_session_id"`
	Orders          []Order_resp_dto `json:"orders"`
	SessionTotal    float64          `json:"session_total"` // Sum of all orders
}

type Order_resp_dto struct {
	ID          uuid.UUID            `json:"id"`
	OrderNumber string               `json:"order_number"`
	Status      string               `json:"status"` // PENDING, SERVED, etc.
	CreatedAt   time.Time            `json:"created_at"`
	Items       []OrderItem_resp_dto `json:"items"`
}

type OrderItem_resp_dto struct {
	ID         uuid.UUID `json:"id"`
	MenuItemID uuid.UUID `json:"menu_item_id"`
	Name       string    `json:"name"` // UI needs to show "Margherita Pizza"
	Quantity   int       `json:"quantity"`
	UnitPrice  float64   `json:"unit_price"` // Price at the time of order
	Status     string    `json:"status"`     // PENDING, COOKING, READY
	Notes      string    `json:"notes"`      // Special instructions

	Variants []Variant_resp_dto `json:"variants"`
	AddOns   []Addon_resp_dto   `json:"addons"`
}

type Variant_resp_dto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`  // e.g., "Extra Large"
	Price float64   `json:"price"` // Price of this specific variant
}

type Addon_resp_dto struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"` // e.g., "Extra Cheese"
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"` // Price per unit of addon
}

/* -------------------------------------------------------------------------- */
/*                          Restaurant General Stuff                          */
/* -------------------------------------------------------------------------- */

type RestaurantGeneralInfo_resp_dto struct {
	RestaurantID            uuid.UUID                                              `json:"restaurant_id"`
	RestaurantLegalEntityID uuid.UUID                                              `json:"resturant_legal_entity_id"`
	Name                    string                                                 `json:"name"`
	OpeningTime             string                                                 `json:"opening_time"`
	ClosingTime             string                                                 `json:"closing_time"`
	IsPureVeg               bool                                                   `json:"is_pure_veg"`
	IsActive                bool                                                   `json:"is_active"`
	IsConfigured            bool                                                   `json:"is_configured"`
	CurrentConfigStatus     restaurant_constants.RestaurantConfigurationStatusType `json:"current_configured_status"`
	LastUpdated             time.Time                                              `json:"last_updated"`
}

type RestaurantPreferredKitchen_resp_dto struct {
	RestaurantID *uuid.UUID `json:"restaurant_id"`
	KitchenID    *uuid.UUID `json:"kitchen_id"`
}

type RestuarantInventoriesResponseDto struct {
	ID                 uuid.UUID `json:"id"`
	InventoryAccountID uuid.UUID `json:"inventory_account_id"`
	BusinessUnitID     uuid.UUID `json:"business_unit_id"`
	LegalEntityID      uuid.UUID `json:"legal_entity_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Location           string    `json:"location"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
