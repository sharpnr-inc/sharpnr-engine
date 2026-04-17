package inventory_engine_contracts

import "github.com/google/uuid"

type UpdateInventoryItemParams struct {
	SupplierID     *uuid.UUID
	COGSAccountID  *uuid.UUID
	AssetAccountID *uuid.UUID
	Name           *string
	Description    *string
	Category       *string
	Unit           *string
	MinStockLevel  *float64
	MaxStockLevel  *float64
	Status         *string
}

// Define the update parameters using pointers for partial updates
type UpdateInventoryParams struct {
	// You can choose to omit the IDs if they shouldn't be changed after creation,
	// but including them allows transferring an inventory between units.
	InventoryAccountID *uuid.UUID
	BusinessUnitID     *uuid.UUID
	LegalEntityID      *uuid.UUID
	Name               *string
	Description        *string
	Location           *string
	Status             *string
}

type UpdateInventoryAccountParams struct {
	Name     *string
	Type     *string
	IsShared *bool
}
