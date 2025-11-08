package accounting_models

import "github.com/google/uuid"

type Customer struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID   uuid.UUID
	Name    string
	Email   string
	Phone   string
	GSTIN   string
	Address string
	Status  string
	Balance float64
}
