package kitchen_contracts

import "github.com/google/uuid"

type CreateKOT_req_dto struct {
	OrganizationID uuid.UUID     `json:"organization_id"`
	PropertyID     uuid.UUID     `json:"property_id"`
	KitchenID      uuid.UUID     `json:"kitchen_id"`
	OrderID        uuid.UUID     `json:"order_id"`
	Items          []KOTItem_dto `json:"items"`
}

type KOTItem_dto struct {
	OrderItemID  uuid.UUID      `json:"order_item_id"` // Link back to RestaurantOrderItem
	MenuItemID   uuid.UUID      `json:"menu_item_id"`
	MenuItemName string         `json:"menu_item_name"`
	ItemName     string         `json:"item_name"`
	Quantity     int            `json:"quantity"`
	Remarks      string         `json:"remarks"`
	Variants     []KOTExtra_dto `json:"variants"`
	Addons       []KOTExtra_dto `json:"addons"`
}

type KOTExtra_dto struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Qty  int       `json:"qty"`
}
