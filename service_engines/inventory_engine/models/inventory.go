package inventory_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Models for inventory account
type InventroyAccount struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `json:"organization_id" gorm:"type:uuid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (InventroyAccount) TableName() string {
	return "inventory.inventory_accounts"
}

type InventoryAccountSettings struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InventoryAccountID uuid.UUID `json:"inventory_account_id" gorm:"type:uuid"`
	IsShared           bool      `json:"is_shared"`
}

func (InventoryAccountSettings) TableName() string {
	return "inventory.inventory_account_settings"
}

type Inventory struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InventoryAccountID uuid.UUID `json:"inventory_account_id" gorm:"type:uuid"`
	BusinessUnitID     uuid.UUID // 👈 important beacuse the this identify the who owns the inventory
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventory.inventories"
}

type IntenventoryItem struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InventoryID uuid.UUID `json:"inventory_id" gorm:"type:uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (IntenventoryItem) TableName() string {
	return "inventory.inventory_items"
}

type InventoryItemPrice struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InventoryItemID uuid.UUID `json:"inventory_item_id" gorm:"type:uuid"`
	Price           float64   `json:"price"`
	Currency        string    `json:"currency"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (InventoryItemPrice) TableName() string {
	return "inventory.inventory_item_prices"
}

type InventoryItemHistory struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InventoryItemID uuid.UUID `json:"inventory_item_id" gorm:"type:uuid"`
	ChangeType      string    `json:"change_type"` // e.g., "addition", "removal", "update"
	QuantityChange  int       `json:"quantity_change"`
	CreatedAt       time.Time `json:"created_at"`
}

func (InventoryItemHistory) TableName() string {
	return "inventory.inventory_item_histories"
}
