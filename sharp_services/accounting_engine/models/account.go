package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

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
	Code string `gorm:"not null"`

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
	Type accountingconstants.AccountGroupType `gorm:"type:varchar(20);not null"`

	// Opening balance for a fiscal year
	OpeningBalance float64 `gorm:"default:0"`
	OpeningIsDebit bool    `gorm:"default:true"`

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
