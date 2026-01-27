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

type AccountingRepositoryInterface interface {

	// ------------------------
	// Accounting Book
	// ------------------------

	CreateAccountingBook(ctx context.Context, accounting_book *accounting_models.AccountingBook) error
	GetAccountingBookByID(ctx context.Context, accounting_bookID uuid.UUID) ([]*accounting_models.AccountingBook, error)
	GetAccountingBookByLegalEntityID(ctx context.Context, legal_entityId uuid.UUID) (*accounting_models.AccountingBook, error)

	// ------------------------
	// Fiscal Year
	// ------------------------

	CreateFiscalYear(ctx context.Context, fy *accounting_models.FiscalYear) error
	GetOpenFiscalYearByBookAndDate(ctx context.Context, bookID uuid.UUID, date time.Time) (*accounting_models.FiscalYear, error)
	CloseFiscalYear(ctx context.Context, fiscalYearID uuid.UUID) error

	// ------------------------
	// Account Groups
	// ------------------------

	CreateAccountGroup(ctx context.Context, group *accounting_models.AccountGroup) error
	GetAccountGroupsByBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.AccountGroup, error)
	GetAccountGroupByID(ctx context.Context, groupID uuid.UUID) (*accounting_models.AccountGroup, error)

	// ------------------------
	// Accounts
	// ------------------------

	CreateAccount(ctx context.Context, account *accounting_models.Account) error
	GetAccountsByGroupID(ctx context.Context, groupID uuid.UUID) ([]accounting_models.Account, error)
	GetAccountByID(ctx context.Context, accountID uuid.UUID) (*accounting_models.Account, error)

	// ------------------------
	// Journal (CORE)
	// ------------------------

	CreateJournalEntry(ctx context.Context, entry *accounting_models.JournalEntry) error
	GetJournalEntryByID(ctx context.Context, entryID uuid.UUID) (*accounting_models.JournalEntry, error)
	GetJournalEntriesByBookAndPeriod(ctx context.Context, bookID uuid.UUID, from, to time.Time) ([]accounting_models.JournalEntry, error)
	GetJournalEntryTotals(ctx context.Context, entryID uuid.UUID) (totalDebit float64, totalCredit float64, err error)

	// ------------------------
	// Journal Lines (Ledger Backbone)
	// ------------------------

	GetJournalLinesByEntryID(ctx context.Context, entryID uuid.UUID) ([]accounting_models.JournalLine, error)
}
