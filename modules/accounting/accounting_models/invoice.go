package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

// TODO : We need to create this or change this into AccountingSoloInvoice
// So when Hotel is other Module is Not there it can work alone

type Invoice struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID       uuid.UUID
	BookID      uuid.UUID
	CustomerID  uuid.UUID
	InvoiceNo   string
	InvoiceDate time.Time
	DueDate     time.Time
	SubTotal    float64
	TaxAmount   float64
	TotalAmount float64
	Status      string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
}
