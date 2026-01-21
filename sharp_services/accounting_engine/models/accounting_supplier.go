package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	Name string

	Type string
	// FOOD, LIQUOR, LAUNDRY, MAINTENANCE, GENERAL

	ContactPerson *string
	Email         *string
	Phone         *string

	GSTIN *string
	PAN   *string

	IsActive bool

	CreatedAt time.Time
}

type SupplierPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID
	AccountingBookID uuid.UUID

	SupplierID uuid.UUID

	PaymentDate time.Time

	Amount float64

	Mode string
	// CASH, BANK, UPI

	Reference string

	CreatedAt time.Time
}
