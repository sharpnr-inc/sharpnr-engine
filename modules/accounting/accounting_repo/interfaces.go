// Package accounting_repo provides the repository layer for all Accounting-related
// data operations. This package acts as the single source of truth for all
// persistence and data retrieval logic within the Accounting module.
package accounting_repo

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
type AccountingRepository interface {
	// Example placeholder methods for guidance
	// (to be implemented in accounting_repo.go or submodules)

	// CreateInvoice persists a new invoice record in the database.
	// Returns the created invoice ID and an error if operation fails.
	// CreateInvoice(ctx context.Context, invoice *accounting_models.Invoice) (uuid.UUID, error)

	// GetLedgerBalance retrieves the ledger balance for a given account and fiscal year.
	// GetLedgerBalance(ctx context.Context, accountID uuid.UUID, fiscalYearID uint) (*accounting_models.LedgerBalance, error)

	// RecordTransaction saves a complete accounting transaction and its entries.
	// RecordTransaction(ctx context.Context, tx *accounting_models.Transaction) error
}
