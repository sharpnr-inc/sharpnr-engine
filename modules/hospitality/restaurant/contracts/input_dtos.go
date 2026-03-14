package restaurant_contracts

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"

	"github.com/google/uuid"
)

type RestaurantFloorLayout_Response_dto struct {
	Floor_Name   string
	Floor_Number int
	Tables       []RestaurantTableLayout_Response_dto
}

type RestaurantLayout_Response_dto struct {
	Layout []RestaurantFloorLayout_Response_dto
}

type RestaurantTableLayout_Response_dto struct {
	TableName            string
	TableCode            string
	FloorID              string
	TableID              string
	Status               restaurant_constants.RestaurantTableStatusType
	Capacity             int
	CurrentDineSessionID *uuid.UUID
}

/* ----------------------------------- *** ---------------------------------- */

// Order represents the top-level order for a table
type DiningSessionOrder_req_dto struct {
	TableID uuid.UUID   `json:"table_id"`
	Items   []OrderItem `json:"items"`
}

// OrderItem matches the "truth" of your React Cart Item
type OrderItem struct {
	MenuItemID          uuid.UUID         `json:"menu_item_id"`
	ItemName            string            `json:"item_name"`
	Quantity            int               `json:"quantity"`
	BasePriceAtOrder    float64           `json:"base_price_at_order"`
	SpecialInstructions string            `json:"special_instructions"`
	SelectedVariants    []SelectedVariant `json:"selected_variants"`
	SelectedAddOns      []SelectedAddOn   `json:"selected_add_ons"`
}

// SelectedVariant handles multiple groups (Size, Crust, etc.)
type SelectedVariant struct {
	VariantGroupID uuid.UUID `json:"variant_group_id"`
	VariantID      uuid.UUID `json:"variant_id"`
	VariantName    string    `json:"variant_name"`
	PriceAtOrder   float64   `json:"price_at_order"`
}

// SelectedAddOn handles multi-select extras (Coke, Extra Cheese)
type SelectedAddOn struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Qty          int       `json:"qty"`
	PriceAtOrder float64   `json:"price_at_order"`
}

type UpdateBodyRequestDto struct {
	Name          *string `json:"name" binding:"omitempty,min=3"`
	OpeningTime   *string `json:"opening_time" binding:"omitempty"`
	ClosingTime   *string `json:"closing_time" binding:"omitempty"`
	LegalEntityID *string `json:"legal_entity_id" binding:"omitempty"`
}

type AddLayoutRequestDto struct {
	RestaurantLayout []SetupRestaurantLayoutInput `json:"restaurant_layout"`
}

/* ---------------------------------- <***> --------------------------------- */
