package accounting_repo

import (
	"context"

	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"

	"time"

	"github.com/google/uuid"
)

type AccountingRepositoryInterface interface {
	AccountingBookRepositoryInterface
	FiscalYearRepositoryInterface
	AccountGroupRepositoryInterface
	AccountOrLedgerRepositoryInterface
	JournalEntryRepositoryInterface
	JournalLineRepositoryInterface
	ReportingRepositoryInterface
	PurchaseInvoiceRepositoryInterface
	EngineAccountRepositoryInterface
}

/* ---------------------------------- <***> --------------------------------- */

// EngineAccountRepositoryInterface owns the org-level account row that says the
// organization has taken the accounting service. It sits above the books: an
// org has one of these, and a book per legal entity under it.
type EngineAccountRepositoryInterface interface {
	InitializeAccountingEngineAccount(ctx context.Context, acc accounting_models.OrgAccountineEngineAccount) error
	CreateSettingForAccountingEngineAccount(ctx context.Context, accSetting accounting_models.AccountingEngineAccountSetting) error
	GetAccountingEngineAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*accounting_models.OrgAccountineEngineAccount, error)

	// AccountingEngineAccountSwitch is the kill switch, and the path a start on
	// an already provisioned-but-disabled org takes: the account is
	// re-activated, never replaced.
	AccountingEngineAccountSwitch(ctx context.Context, orgID, orgAccountingAccountID uuid.UUID, isActive bool) error
}

/* ---------------------------------- <***> --------------------------------- */

type AccountingBookRepositoryInterface interface {
	CreateAccountingBook(ctx context.Context, accounting_book *accounting_models.AccountingBook) error
	GetAccountingBookByID(ctx context.Context, accounting_bookID uuid.UUID) (*accounting_models.AccountingBook, error)
	GetAccountingBookByLegalEntityID(ctx context.Context, legal_entityId uuid.UUID) (*accounting_models.AccountingBook, error)
}

/* ---------------------------------- <***> --------------------------------- */
type FiscalYearRepositoryInterface interface {
	CreateFiscalYear(ctx context.Context, fy *accounting_models.FiscalYear) error
	GetAllFiscalYearsByBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.FiscalYear, error)
	GetOpenFiscalYearByBookAndDate(ctx context.Context, bookID uuid.UUID, date time.Time) (*accounting_models.FiscalYear, error)
	CloseFiscalYear(ctx context.Context, fiscalYearID uuid.UUID) error
}

/* ---------------------------------- <***> --------------------------------- */
type AccountGroupRepositoryInterface interface {
	CreateAccountGroup(ctx context.Context, group *accounting_models.AccountGroup) error
	GetAccountGroupsByBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.AccountGroup, error)
	GetAccountGroupByID(ctx context.Context, groupID uuid.UUID) (*accounting_models.AccountGroup, error)
}

/* ---------------------------------- <***> --------------------------------- */
type AccountOrLedgerRepositoryInterface interface {
	CreateAccount(ctx context.Context, account *accounting_models.Account) error
	GetAccountsByGroupID(ctx context.Context, groupID uuid.UUID) ([]accounting_models.Account, error)
	GetAccountByID(ctx context.Context, accountID uuid.UUID) (*accounting_models.Account, error)
	GetLedger(accountID uuid.UUID, from, to time.Time) ([]accountingdtos.LedgerLineDto, error)
}

/* ---------------------------------- <***> --------------------------------- */
type JournalEntryRepositoryInterface interface {
	CreateJournalEntry(ctx context.Context, entry *accounting_models.JournalEntry) error
	GetJournalEntryByID(ctx context.Context, entryID uuid.UUID) (*accounting_models.JournalEntry, error)
	GetJournalEntriesByAccountingBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.JournalEntry, error)
	GetJournalEntriesByBookAndPeriod(ctx context.Context, bookID uuid.UUID, from, to time.Time) ([]accounting_models.JournalEntry, error)
	GetJournalEntryTotals(ctx context.Context, entryID uuid.UUID) (totalDebit float64, totalCredit float64, err error)
}

/* ---------------------------------- <***> --------------------------------- */
type JournalLineRepositoryInterface interface {
	CreateJournalLine(ctx context.Context, line *accounting_models.JournalLine) error
	GetJournalLinesByJournalEntryID(ctx context.Context, entryID uuid.UUID) ([]accounting_models.JournalLine, error)
	GetJournalLinesByAccountID(ctx context.Context, accountID uuid.UUID) ([]accounting_models.JournalLine, error)
}

/* ---------------------------------- <***> --------------------------------- */

type ReportingRepositoryInterface interface {
	GetAccountMovementsForPeriod(
		ctx context.Context,
		bookID uuid.UUID,
		from, to time.Time,
	) ([]accountingdtos.AccountMovementRow, error)
}

type PurchaseInvoiceRepositoryInterface interface {

	// ── Purchase Invoice ──────────────────────────────────────────────────────

	CreatePurchaseInvoice(ctx context.Context, invoice *accounting_models.PurchaseInvoice) error

	GetPurchaseInvoiceByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchaseInvoice, error)

	// for duplicate check — InvoiceNumber must be unique per Supplier
	GetPurchaseInvoiceBySupplierAndInvoiceNumber(ctx context.Context, supplierID uuid.UUID, invoiceNumber string) (*accounting_models.PurchaseInvoice, error)

	ListPurchaseInvoices(ctx context.Context, filter PurchaseInvoiceFilter) ([]*accounting_models.PurchaseInvoice, int64, error)
	// int64 = total count for pagination

	// UpdatePurchaseInvoice(ctx context.Context, invoice *accounting_models.PurchaseInvoice) error

	UpdatePurchaseInvoiceStatus(ctx context.Context, id uuid.UUID, status accountingconstants.AccountingPurchaseInvoiceStatus) error

	UpdatePurchaseInvoicePaymentStatus(ctx context.Context, id uuid.UUID, paymentStatus string, amountPaid float64, amountDue float64) error

	// set when journal entry is auto-created on POST
	SetPurchaseInvoiceJournalEntry(ctx context.Context, id uuid.UUID, journalEntryID uuid.UUID) error

	// DeletePurchaseInvoice(ctx context.Context, id uuid.UUID) error
	// soft-delete or status → CANCELLED

	// ── Purchase Invoice Lines ────────────────────────────────────────────────

	CreatePurchaseInvoiceLine(ctx context.Context, line *accounting_models.PurchaseInvoiceLine) error

	CreatePurchaseInvoiceLines(ctx context.Context, lines []*accounting_models.PurchaseInvoiceLine) error
	// batch insert — always prefer this over looping CreatePurchaseInvoiceLine

	GetPurchaseInvoiceLineByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchaseInvoiceLine, error)

	ListPurchaseInvoiceLines(ctx context.Context, purchaseInvoiceID uuid.UUID) ([]*accounting_models.PurchaseInvoiceLine, error)

	// UpdatePurchaseInvoiceLine(ctx context.Context, line *accounting_models.PurchaseInvoiceLine) error

	// DeletePurchaseInvoiceLine(ctx context.Context, id uuid.UUID) error

	// DeletePurchaseInvoiceLinesByInvoiceID(ctx context.Context, purchaseInvoiceID uuid.UUID) error
	// used when replacing all lines on an edit

	// ── Purchase Payments ─────────────────────────────────────────────────────

	CreatePurchasePayment(ctx context.Context, payment *accounting_models.PurchasePayment) error

	GetPurchasePaymentByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchasePayment, error)

	ListPurchasePaymentsByInvoice(ctx context.Context, purchaseInvoiceID uuid.UUID) ([]*accounting_models.PurchasePayment, error)

	ListPurchasePaymentsBySupplier(ctx context.Context, supplierID uuid.UUID) ([]*accounting_models.PurchasePayment, error)

	// set after journal entry auto-created for the payment
	// SetPurchasePaymentJournalEntry(ctx context.Context, id uuid.UUID, journalEntryID uuid.UUID) error

	// DeletePurchasePayment(ctx context.Context, id uuid.UUID) error

	// ── Aggregates / Reporting ────────────────────────────────────────────────

	GetSupplierOutstandingBalance(ctx context.Context, supplierID uuid.UUID, accountingBookID uuid.UUID) (float64, error)
	// SUM(amount_due) WHERE supplier_id = ? AND payment_status != 'PAID'

	GetPurchaseInvoiceSummary(ctx context.Context, accountingBookID uuid.UUID, fiscalYearID uuid.UUID) (*PurchaseInvoiceSummary, error)
	// total invoiced, total paid, total outstanding — for dashboard

	ListOverduePurchaseInvoices(ctx context.Context, accountingBookID uuid.UUID, asOf time.Time) ([]*accounting_models.PurchaseInvoice, error)
	// WHERE due_date < asOf AND payment_status != 'PAID'
}

// ── Filter struct ─────────────────────────────────────────────────────────────

type PurchaseInvoiceFilter struct {
	AccountingBookID *uuid.UUID
	FiscalYearID     *uuid.UUID
	SupplierID       *uuid.UUID

	Status        *accountingconstants.AccountingPurchaseInvoiceStatus
	PaymentStatus *string // UNPAID | PARTIAL | PAID

	// date range on InvoiceDate
	FromDate *time.Time
	ToDate   *time.Time

	// search across InvoiceNumber + OurReference
	Search *string

	// pagination
	Page  int // 1-based
	Limit int // default 20
}

// ── Summary DTO ───────────────────────────────────────────────────────────────

type PurchaseInvoiceSummary struct {
	TotalInvoiced    float64
	TotalPaid        float64
	TotalOutstanding float64
	TotalOverdue     float64
	InvoiceCount     int64
}
