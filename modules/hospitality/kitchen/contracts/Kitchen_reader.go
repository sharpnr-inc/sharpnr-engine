package kitchen_contracts

import "github.com/google/uuid"

type Kitchen_Reader struct {
	ID          uuid.UUID
	Name        string
	Description string
	Status      bool
	OrgID       uuid.UUID
	PropertyID  uuid.UUID
}
