package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

// AccountingBook represents a complete and independent set of accounting records
// within an organization.
//
// An organization (tenant) may have multiple accounting books to maintain
// separate financial records for different business units or departments,
// such as Hotel, Restaurant, Bar, Factory, etc.
//
// Each AccountingBook:
//   - Belongs to exactly one Organization (OrgID)
//   - Has its own Fiscal Years
//   - Owns its own Account Groups, Accounts, and Journal Entries
//   - Is fully isolated from other accounting books of the same organization
//
// All accounting operations (journal entries, ledgers, trial balance,
// profit & loss, balance sheet) are always scoped to a single AccountingBook.
type AccountingBook struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// OrgID is the owning Organization's identifier.
	//
	// This acts as the tenant identifier for the accounting book.
	// All accounting data under this book is isolated to this organization.
	OrgID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Name is a human-readable name of the accounting book.
	//
	// Examples:
	//   - "Hotel Accounting"
	//   - "Restaurant Accounts"
	//   - "Bar Operations"
	Name string `gorm:"not null"`

	// Department optionally identifies the internal department
	// this accounting book belongs to.
	//
	// This is purely informational and does not affect accounting logic.
	Department *string

	// Module optionally identifies the application module
	// or business domain using this accounting book.
	//
	// Examples:
	//   - "HOSPITALITY"
	//   - "MANUFACTURING"
	//   - "RETAIL"
	Module *string

	// FiscalYears contains all fiscal years configured for this accounting book.
	//
	// At any given time, only one FiscalYear should be in OPEN state.
	FiscalYears []FiscalYear `gorm:"foreignKey:AccountingBookID"`

	// AccountGroups contains the chart of accounts grouped
	// under this accounting book.
	//
	// Each AccountGroup further contains multiple Accounts
	// which are used in journal entries.
	AccountGroups []AccountGroup `gorm:"foreignKey:AccountingBookID"`

	// Status represents the current lifecycle state of the accounting book.
	//
	// Example values:
	//   - ACTIVE
	//   - ARCHIVED
	//   - CLOSED
	Status accountingconstants.AccountingBookStatus

	// CreatedAt is the timestamp when the accounting book was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the accounting book was last updated.
	UpdatedAt time.Time
}
