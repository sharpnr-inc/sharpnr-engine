package inventory_dtos

import (
	inventory_constants "sharpnr-engine/service_engines/inventory_engine/constants"
	"time"

	"github.com/google/uuid"
)

// NewInventoryItemRequestDto
type NewInventoryItemRequestDto struct {
	SupplierID     *uuid.UUID `json:"supplier_id"`
	COGSAccountID  *uuid.UUID `json:"cogs_account_id"`
	AssetAccountID *uuid.UUID `json:"asset_account_id"`
	Name           string     `json:"name" validate:"required"`
	SKU            string     `json:"sku" validate:"required"`
	Description    string     `json:"description"`
	Category       string     `json:"category" validate:"required,oneof=FOOD BEVERAGE PACKAGING OTHER"`
	Unit           string     `json:"unit" validate:"required,oneof=KG LITRE PIECE BOX DOZEN"`
	MinStockLevel  float64    `json:"min_stock_level"`
	MaxStockLevel  float64    `json:"max_stock_level"`
}

// UpdateInventoryItemRequestDto
type UpdateInventoryItemRequestDto struct {
	Name           *string    `json:"name"`
	Description    *string    `json:"description"`
	Category       *string    `json:"category"`
	Unit           *string    `json:"unit"`
	MinStockLevel  *float64   `json:"min_stock_level"`
	MaxStockLevel  *float64   `json:"max_stock_level"`
	SupplierID     *uuid.UUID `json:"supplier_id"`
	COGSAccountID  *uuid.UUID `json:"cogs_account_id"`
	AssetAccountID *uuid.UUID `json:"asset_account_id"`
	Status         *string    `json:"status"`
}

// SetInventoryItemPriceRequestDto
type SetInventoryItemPriceRequestDto struct {
	SupplierID    *uuid.UUID                                 `json:"supplier_id"`
	PriceType     inventory_constants.InventoryItemPriceType `json:"price_type" validate:"required"`
	Price         float64                                    `json:"price" validate:"required,gt=0"`
	Currency      string                                     `json:"currency" validate:"required"`
	EffectiveFrom time.Time                                  `json:"effective_from" validate:"required"`
}
