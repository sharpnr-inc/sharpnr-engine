package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountingBookID uuid.UUID `gorm:"type:uuid;index"`          // newly added
	AccountGroupID   uuid.UUID `gorm:"type:uuid;index"`          // Accounts are grouped under Account Groups (e.g. Cash, Bank, Receivables under Assets group)
	Name             string    `gorm:"not null"`                 // Name of the account (e.g. Cash, Bank, Accounts Receivable)
	Code             string    `gorm:"not null"`                 // Unique code for the account (e.g. 1001 for Cash, 1002 for Bank, etc.)
	OpeningBalance   float64   `gorm:"default:0"`                // Balance at the beginning of the fiscal year
	OpeningIsDebit   bool      `gorm:"default:true"`             // true if opening balance is debit, false if credit
	Status           string    `gorm:"size:20;default:'ACTIVE'"` // ACTIVE, INACTIVE, CLOSED
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Account) TableName() string {
	return "accounting_engine.account"
}
