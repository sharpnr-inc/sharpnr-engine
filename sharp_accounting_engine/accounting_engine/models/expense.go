package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID `gorm:"type:uuid;not null;index"`
	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	FiscalYearID     uuid.UUID `gorm:"type:uuid;not null;index"`

	ExpenseDate time.Time `gorm:"not null"`

	// Category → FK instead of raw string
	ExpenseCategoryID *uuid.UUID `gorm:"type:uuid;index"`
	// fallback if no category FK
	CategoryLabel string `gorm:"size:100"`

	Description string `gorm:"not null"`

	// Payee (optional)
	PayeeName  string
	SupplierID *uuid.UUID `gorm:"type:uuid;index"` // if payee is a known supplier

	// Amounts
	SubTotal    float64 `gorm:"default:0"`
	TaxAmount   float64 `gorm:"default:0"`
	TotalAmount float64 `gorm:"default:0"`

	TaxCategoryID *uuid.UUID `gorm:"type:uuid"`

	// Accounting linkage
	DebitAccountID  uuid.UUID  `gorm:"type:uuid;not null"` // Expense a/c
	CreditAccountID uuid.UUID  `gorm:"type:uuid;not null"` // Cash / Bank a/c
	JournalEntryID  *uuid.UUID `gorm:"type:uuid"`          // set on POST

	// Payment
	PaymentMode string `gorm:"size:20"` // CASH | UPI | BANK | CHEQUE
	Reference   string // UTR, cheque no, etc.

	// Status
	Status string `gorm:"size:20;default:'DRAFT'"` // DRAFT | POSTED | VOID

	// Attachment
	AttachmentURL string

	Narration string
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Expense) TableName() string {
	return "accounting_engine.expense"
}
