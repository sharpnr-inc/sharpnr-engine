package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

// Account represents an individual ledger account within an AccountGroup.
//
// An Account is the smallest and most important unit in the accounting system.
// All financial movements are ultimately recorded against Accounts via
// JournalLines.
//
// Examples of Accounts:
//   - Cash
//   - Bank - HDFC
//   - Accounts Receivable
//   - Room Revenue - Deluxe
//   - Food Revenue
//   - Electricity Expense
//
// Each Account:
//   - Belongs to exactly one AccountGroup
//   - Inherits its AccountType from the parent AccountGroup
//   - Is referenced directly by JournalLines
//
// Accounts should be treated as stable entities; renaming is allowed,
// but changing the accounting nature (type) is not.
type Account struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Name is the human-readable name of the account.
	//
	// Example:
	//   - "Cash"
	//   - "Room Revenue"
	//   - "Electricity Expense"
	Name string `gorm:"not null"`

	// Code is a unique identifier for the account within the system.
	//
	// This is typically used for:
	//   - Accounting standards
	//   - Reporting
	//   - External integrations
	//
	// Example:
	//   - "1001" (Cash)
	//   - "4001" (Room Revenue)
	Code string `gorm:"uniqueIndex;not null"`

	// AccountGroupID is the identifier of the account group
	// this account belongs to.
	AccountGroupID uuid.UUID `gorm:"type:uuid;index"`

	// AccountGroup is the parent account group.
	// Used for ORM relationships and preloading.
	AccountGroup AccountGroup `gorm:"foreignKey:AccountGroupID"`

	// Type defines the accounting nature of this account.
	//
	// This value must always match the Type of the parent AccountGroup.
	// It should never be changed after creation.
	Type accountingconstants.AccountType `gorm:"type:varchar(20);not null"`

	// IsDefault indicates whether this is a system-defined default account.
	//
	// Default accounts are created during initial accounting setup
	// and should typically be protected from deletion.
	IsDefault bool `gorm:"default:false"`

	// Status represents the current state of the account.
	//
	// Example values:
	//   - ACTIVE
	//   - DISABLED
	//   - ARCHIVED
	Status string `gorm:"size:20;default:'ACTIVE'"`

	// CreatedAt is the timestamp when the account was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the account was last updated.
	UpdatedAt time.Time
}
