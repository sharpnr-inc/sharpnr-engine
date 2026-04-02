package inventory_engine_contracts

import "github.com/google/uuid"

type CreateInventoryInInvetoryAccountInternalContract struct {
	InventoryAccountID uuid.UUID
	BusinessUnitID     uuid.UUID
	LegalEntityID      uuid.UUID
	Name               string
	Description        string
	Location           string
}
