package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Bill struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID       uuid.UUID
	BookID      uuid.UUID
	SupplierID  uuid.UUID
	BillNo      string
	BillDate    time.Time
	DueDate     time.Time
	SubTotal    float64
	TaxAmount   float64
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
}
