package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

// AccountGroup represents a logical classification of accounts
// within an AccountingBook.
//
// Account groups form the "Chart of Accounts" structure and are used
// to organize individual accounts under standard accounting categories
// such as Assets, Liabilities, Income, Expenses, and Equity.
//
// Examples of AccountGroups:
//   - Cash & Bank (ASSET)
//   - Accounts Receivable (ASSET)
//   - Room Revenue (INCOME)
//   - Food Revenue (INCOME)
//   - Operating Expenses (EXPENSE)
//   - Capital & Reserves (EQUITY)
//
// Each AccountGroup:
//   - Belongs to exactly one AccountingBook
//   - Has a fixed AccountType that defines its accounting nature
//   - Contains multiple Accounts used in journal entries
//
// The AccountType of an AccountGroup determines how balances behave
// in reports such as Trial Balance, Profit & Loss, and Balance Sheet.
type AccountGroup struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Name is the human-readable name of the account group.
	//
	// Example:
	//   - "Cash & Bank"
	//   - "Room Revenue"
	//   - "Operating Expenses"
	Name string `gorm:"not null"`

	// AccountingBookID is the identifier of the accounting book
	// this account group belongs to.
	//
	// All accounts under this group are scoped to the same accounting book.
	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`

	// AccountingBook is the parent accounting book.
	// This field is used for ORM relationships and preloading.
	AccountingBook AccountingBook `gorm:"foreignKey:AccountingBookID"`

	// Accounts contains the individual ledger accounts
	// that belong to this account group.
	//
	// Journal entries always reference Accounts, never AccountGroups directly.
	Accounts []Account `gorm:"foreignKey:AccountGroupID"`

	// Type defines the accounting nature of this group.
	//
	// Allowed values:
	//   - ASSET
	//   - LIABILITY
	//   - EQUITY
	//   - INCOME
	//   - EXPENSE
	//
	// This value must remain immutable once the group is created.
	Type accountingconstants.AccountType `gorm:"type:varchar(20);not null"`

	// CreatedAt is the timestamp when the account group was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the account group was last updated.
	UpdatedAt time.Time
}
