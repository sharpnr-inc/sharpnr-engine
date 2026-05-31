package inventory_engine_models

import (
	inventory_constants "sharpnr-engine/sharp_inventory_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Models for inventory account
type InventoryAccount struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name           string    `gorm:"not null"`      // "Main Warehouse", "Kitchen Store"
	Type           string    `gorm:"size:20"`       // WAREHOUSE | KITCHEN | STORE
	IsShared       bool      `gorm:"default:false"` // shared across business units?
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (InventoryAccount) TableName() string {
	return "inventory.inventory_accounts"
}

// ── Inventory (per business unit) ──
type Inventory struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	InventoryAccountID uuid.UUID `gorm:"type:uuid;not null;index"`
	BusinessUnitID     uuid.UUID `gorm:"type:uuid;not null;index"` // Restaurant / Hotel
	LegalEntityID      uuid.UUID `gorm:"type:uuid;not null;index"` // for accounting linkage
	Name               string    `gorm:"not null"`
	Description        string
	Location           string // "Ground Floor Kitchen", "Cold Storage"
	Status             string `gorm:"size:20;default:'ACTIVE'"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (Inventory) TableName() string {
	return "inventory.inventories"
}

// ── Inventory Item ──
type InventoryItem struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey"`
	InventoryID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	SupplierID     *uuid.UUID `gorm:"type:uuid;index"` // preferred supplier
	COGSAccountID  *uuid.UUID `gorm:"type:uuid"`       // accounting linkage ← KEY
	AssetAccountID *uuid.UUID `gorm:"type:uuid"`       // if held as asset

	Name        string `gorm:"not null"`
	SKU         string `gorm:"uniqueIndex"` // e.g. "VEG-CHKN-001"
	Description string
	Category    string `gorm:"size:30"` // FOOD | BEVERAGE | PACKAGING | OTHER
	Unit        string `gorm:"size:20"` // KG | LITRE | PIECE | BOX | DOZEN

	Quantity      float64 `gorm:"default:0"` // float for KG/LITRE
	MinStockLevel float64 `gorm:"default:0"` // alert threshold
	MaxStockLevel float64 `gorm:"default:0"` // for over-ordering prevention

	Status    string `gorm:"size:20;default:'ACTIVE'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (InventoryItem) TableName() string {
	return "inventory.inventory_items"
}

// ── Item Price (per supplier, with history) ──
type InventoryItemPrice struct {
	ID              uuid.UUID                                  `gorm:"type:uuid;primaryKey"`
	InventoryItemID uuid.UUID                                  `gorm:"type:uuid;not null;index"`
	SupplierID      *uuid.UUID                                 `gorm:"type:uuid;index"`
	PriceType       inventory_constants.InventoryItemPriceType `gorm:"size:20"` // PURCHASE_PRICE | SELLING_PRICE
	Price           float64                                    `gorm:"not null"`
	Currency        string                                     `gorm:"size:5;default:'INR'"`
	EffectiveFrom   time.Time                                  `gorm:"not null"`
	EffectiveTo     *time.Time                                 // null = currently active
	CreatedAt       time.Time
}

func (InventoryItemPrice) TableName() string {
	return "inventory.inventory_item_prices"
}

// ── Item History (every stock movement) ──
type InventoryItemHistory struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	InventoryItemID uuid.UUID `gorm:"type:uuid;not null;index"`

	ChangeType string `gorm:"size:20;not null"`
	// PURCHASE_RECEIVED | CONSUMED | ADJUSTED | WASTAGE | TRANSFERRED | RETURNED

	ReferenceType string     `gorm:"size:20"`   // PO | CONSUMPTION | ADJUSTMENT
	ReferenceID   *uuid.UUID `gorm:"type:uuid"` // links to PO or consumption record

	QuantityBefore float64
	QuantityChange float64 // positive = IN, negative = OUT
	QuantityAfter  float64

	UnitCost  float64 // cost per unit at time of transaction ← for COGS calc
	TotalCost float64 // QuantityChange * UnitCost

	PerformedBy uuid.UUID `gorm:"type:uuid"` // employee ID
	Notes       string

	CreatedAt time.Time
}

func (InventoryItemHistory) TableName() string {
	return "inventory.inventory_item_histories"
}
