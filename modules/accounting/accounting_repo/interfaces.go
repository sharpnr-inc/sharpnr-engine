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
type AccountingRepositoryInterface interface {
	WithDbTrnx()
}
