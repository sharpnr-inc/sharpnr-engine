package inventory_engine_contracts

import "github.com/google/uuid"

type InventoryItemsListOutputContracts struct {
	ItemID      uuid.UUID
	InventoryID uuid.UUID
	SuppliersID *uuid.UUID

	Name        string
	SKU         string
	Description string
	Unit        string

	Quantity      float64
	Status        string
	PurchasePrice float64
	SellingPrice  float64
	
}
