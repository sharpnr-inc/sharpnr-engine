package supplier_contracts

import "github.com/google/uuid"

type SuppliersOutput struct {
	ID            uuid.UUID
	SuppliersName string
}
