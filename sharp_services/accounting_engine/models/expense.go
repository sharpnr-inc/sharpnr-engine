package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID
	AccountingBookID uuid.UUID

	ExpenseDate time.Time

	Category string
	// ELECTRICITY, FUEL, REPAIR, MARKETING

	Description string

	Amount float64

	PaymentMode string
	// CASH, UPI, BANK

	Reference string

	CreatedBy uuid.UUID

	CreatedAt time.Time
}

func (Expense) TableName() string {
	return "accounting_engine.expense"
}

type PurchaseBill struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID
	AccountingBookID uuid.UUID

	SupplierID uuid.UUID

	BillNumber string
	BillDate   time.Time
	DueDate    *time.Time

	SubTotal  float64
	TaxAmount float64
	Total     float64

	Status string
	// DRAFT, POSTED, PAID, CANCELLED

	CreatedAt time.Time
}

func (PurchaseBill) TableName() string {
	return "accounting_engine.purchase_bill"
}

type PurchaseBillItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	PurchaseBillID uuid.UUID

	Description string

	Quantity float64
	Rate     float64

	Amount float64

	ExpenseAccountCode string
}

func (PurchaseBillItem) TableName() string {
	return "accounting_engine.purchase_bill_item"
}
