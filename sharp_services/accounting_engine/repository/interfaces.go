// Package accounting_repo provides the repository layer for all Accounting-related
// data operations. This package acts as the single source of truth for all
// persistence and data retrieval logic within the Accounting module.
package accounting_repo

import (
	"context"
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"

	"time"

	"github.com/google/uuid"
)

// AccountingRepository defines the core contract for data access and persistence
// operations in the Accounting module.
//
// This interface acts as the “Single Point of Truth” for accounting data flow,
// ensuring all read/write operations for transactions, invoices, bills, ledgers,
// and financial summaries are abstracted from the service layer.
//
// Responsibilities include, but are not limited to:
//   - Managing persistence of all accounting entities (Accounts, Ledgers, Invoices, etc.)
//   - Providing transactional consistency across modules (e.g., Bar, Hotel, HR)
//   - Serving as a stable data abstraction for services, controllers, and event handlers
//
// Example:
//
//	repo := NewAccountingRepository(db)
//	ledger, err := repo.GetLedgerByAccountID(accountID)
//
// Implementations of this interface should use GORM or any other
// database ORM/driver under the hood for persistence.
type AccountingRepositoryInterface interface {

	// ------------------------
	// Transaction handling
	// ------------------------
	WithTx(
		ctx context.Context,
		fn func(txRepo AccountingRepositoryInterface) error,
	) error

	// ------------------------
	// Accounting Book
	// ------------------------
	CreateAccountingBook(
		ctx context.Context,
		book *accounting_models.AccountingBook,
	) error

	GetAccountingBookByID(
		ctx context.Context,
		bookID uuid.UUID,
	) (*accounting_models.AccountingBook, error)

	GetAccountingBooksByOrgID(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]accounting_models.AccountingBook, error)

	// GetAccountingBookByOrgAndSubModule fetches an accounting book
	// scoped to an organization and a specific sub-module.
	//
	// Returns:
	//   - (*AccountingBook, nil) if found
	//   - (nil, nil) if not found
	//   - (nil, error) on database errors
	GetAccountingBookByOrgAndSubModule(
		ctx context.Context,
		orgID uuid.UUID,
		subModule string,
	) (*accounting_models.AccountingBook, error)

	// GetAccountingBookByOrgAndModule fetches an accounting book
	// scoped to an organization and a module.
	//
	// Returns:
	//   - (*AccountingBook, nil) if found
	//   - (nil, nil) if not found
	//   - (nil, error) on database errors
	GetAccountingBookByOrgAndModule(
		ctx context.Context,
		orgID uuid.UUID,
		module string,
	) (*accounting_models.AccountingBook, error)

	// ------------------------
	// Fiscal Year
	// ------------------------
	CreateFiscalYear(
		ctx context.Context,
		fy *accounting_models.FiscalYear,
	) error

	GetOpenFiscalYearByBookAndDate(
		ctx context.Context,
		bookID uuid.UUID,
		date time.Time,
	) (*accounting_models.FiscalYear, error)

	CloseFiscalYear(
		ctx context.Context,
		fiscalYearID uuid.UUID,
	) error

	// ------------------------
	// Account Groups
	// ------------------------
	CreateAccountGroup(
		ctx context.Context,
		group *accounting_models.AccountGroup,
	) error

	GetAccountGroupsByBookID(
		ctx context.Context,
		bookID uuid.UUID,
	) ([]accounting_models.AccountGroup, error)

	GetAccountGroupByID(
		ctx context.Context,
		groupID uuid.UUID,
	) (*accounting_models.AccountGroup, error)

	// ------------------------
	// Accounts
	// ------------------------
	CreateAccount(
		ctx context.Context,
		account *accounting_models.Account,
	) error

	GetAccountsByGroupID(
		ctx context.Context,
		groupID uuid.UUID,
	) ([]accounting_models.Account, error)

	GetAccountByID(
		ctx context.Context,
		accountID uuid.UUID,
	) (*accounting_models.Account, error)

	// ------------------------
	// Journal (CORE)
	// ------------------------
	CreateJournalEntry(
		ctx context.Context,
		entry *accounting_models.JournalEntry,
	) error

	GetJournalEntryByID(
		ctx context.Context,
		entryID uuid.UUID,
	) (*accounting_models.JournalEntry, error)

	GetJournalEntriesByBookAndPeriod(
		ctx context.Context,
		bookID uuid.UUID,
		from, to time.Time,
	) ([]accounting_models.JournalEntry, error)

	GetJournalEntries(
		ctx context.Context,
		bookID uuid.UUID,
		from time.Time,
		to time.Time,
		limit int,
		offset int,
	) ([]accounting_models.JournalEntry, int64, error)

	GetJournalEntryTotals(
		ctx context.Context,
		entryID uuid.UUID,
	) (totalDebit float64, totalCredit float64, err error)

	GetJournalEntryWithLines(
		ctx context.Context,
		entryID uuid.UUID,
	) (*accounting_models.JournalEntry, []accounting_models.JournalLine, error)

	// ------------------------
	// Journal Lines (Ledger Backbone)
	// ------------------------

	// Batch insert of journal lines (used during posting)
	CreateJournalLines(
		ctx context.Context,
		lines []accounting_models.JournalLine,
	) error

	// Fetch all lines for a given journal entry (voucher view)
	GetJournalLinesByEntryID(
		ctx context.Context,
		entryID uuid.UUID,
	) ([]accounting_models.JournalLine, error)

	// Fetch lines for a given account within a period (ledger, trial balance)
	GetJournalLinesByAccountAndPeriod(
		ctx context.Context,
		accountID uuid.UUID,
		from, to time.Time,
	) ([]accounting_models.JournalLine, error)

	/***************************************************
	 * ⭐ Account Mapping *
	***************************************************/

	// This repo Method return array of Accounting Mapping in Array
	// it also need sub-module of module like BAR, HOTEL in Hospitality module
	GetAccountMappingsByOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]accounting_models.AccountMapping, error)

	CreateAccountMapping(
		ctx context.Context,
		accountMapping *accounting_models.AccountMapping,
	) error

	ResolveAccountMapping(
		ctx context.Context,
		orgID uuid.UUID,
		module string,
		sourceType string,
		sourceCode string,
	) (*accounting_models.Account, error)
}
