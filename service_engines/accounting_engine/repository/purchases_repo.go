package accounting_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *accountingRepository) CreatePurchaseInvoice(ctx context.Context, invoice *accounting_models.PurchaseInvoice) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Use WithContext to link the database operation to the request lifecycle
	if err := db.WithContext(ctx).Create(invoice).Error; err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetPurchaseInvoiceByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchaseInvoice, error) {
	var invoice accounting_models.PurchaseInvoice

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Execute the query with the provided context
	err := db.WithContext(ctx).Where("id = ?", id).First(&invoice).Error

	if err != nil {
		// Return nil and the error (e.g., gorm.ErrRecordNotFound)
		// so the caller can handle 404 vs 500 errors
		return nil, err
	}

	return &invoice, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetPurchaseInvoiceBySupplierAndInvoiceNumber(ctx context.Context, supplierID uuid.UUID, invoiceNumber string) (*accounting_models.PurchaseInvoice, error) {
	var invoice accounting_models.PurchaseInvoice

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Querying with both supplier_id and invoice_number
	err := db.WithContext(ctx).
		Where("supplier_id = ? AND invoice_number = ?", supplierID, invoiceNumber).
		First(&invoice).Error

	if err != nil {
		// Returns gorm.ErrRecordNotFound if no duplicate exists
		return nil, err
	}

	return &invoice, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) ListPurchaseInvoices(ctx context.Context, filter PurchaseInvoiceFilter) ([]*accounting_models.PurchaseInvoice, int64, error) {
	var invoices []*accounting_models.PurchaseInvoice
	var total int64

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Initialize query on the model
	query := db.WithContext(ctx).Model(&accounting_models.PurchaseInvoice{})

	// 1. Apply Functional Filters
	if filter.AccountingBookID != nil {
		query = query.Where("accounting_book_id = ?", filter.AccountingBookID)
	}
	if filter.FiscalYearID != nil {
		query = query.Where("fiscal_year_id = ?", filter.FiscalYearID)
	}
	if filter.SupplierID != nil {
		query = query.Where("supplier_id = ?", filter.SupplierID)
	}
	if filter.Status != nil {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.PaymentStatus != nil {
		query = query.Where("payment_status = ?", filter.PaymentStatus)
	}

	// 2. Apply Date Range Filters
	if filter.FromDate != nil {
		query = query.Where("invoice_date >= ?", filter.FromDate)
	}
	if filter.ToDate != nil {
		query = query.Where("invoice_date <= ?", filter.ToDate)
	}

	// 3. Apply Search (Case-insensitive partial match)
	if filter.Search != nil && *filter.Search != "" {
		searchPattern := "%" + *filter.Search + "%"
		// Using ILIKE for PostgreSQL; use LIKE for MySQL
		query = query.Where("invoice_number ILIKE ? OR our_reference ILIKE ?", searchPattern, searchPattern)
	}

	// 4. Get Total Count (before pagination)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 5. Apply Pagination Defaults
	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	// 6. Execute Final Query
	// We order by date descending so the newest invoices appear first
	err := query.Limit(limit).
		Offset(offset).
		Order("invoice_date DESC, created_at DESC").
		Find(&invoices).Error

	if err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) UpdatePurchaseInvoiceStatus(ctx context.Context, id uuid.UUID, status accountingconstants.AccountingPurchaseInvoiceStatus) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We use .Model() to tell GORM which table/struct we are targeting
	result := db.WithContext(ctx).
		Model(&accounting_models.PurchaseInvoice{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	// Optional: verify that a record was actually found and updated
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) UpdatePurchaseInvoicePaymentStatus(
	ctx context.Context,
	id uuid.UUID,
	paymentStatus string,
	amountPaid float64,
	amountDue float64,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We use a map to ensure that fields with 0 values (like amount_due)
	// are explicitly updated in the database.
	result := db.WithContext(ctx).
		Model(&accounting_models.PurchaseInvoice{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"payment_status": paymentStatus,
			"amount_paid":    amountPaid,
			"amount_due":     amountDue,
			"updated_at":     time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                         Purchase Line Repo Methods                         */
/* -------------------------------------------------------------------------- */

func (r *accountingRepository) CreatePurchaseInvoiceLine(ctx context.Context, line *accounting_models.PurchaseInvoiceLine) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	if err := db.WithContext(ctx).Create(line).Error; err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) CreatePurchaseInvoiceLines(ctx context.Context, lines []*accounting_models.PurchaseInvoiceLine) error {
	if len(lines) == 0 {
		return nil
	}

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// GORM creates all records in a single batch insert
	if err := db.WithContext(ctx).Create(lines).Error; err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetPurchaseInvoiceLineByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchaseInvoiceLine, error) {
	var line accounting_models.PurchaseInvoiceLine

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Where("id = ?", id).First(&line).Error
	if err != nil {
		return nil, err
	}

	return &line, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) ListPurchaseInvoiceLines(ctx context.Context, purchaseInvoiceID uuid.UUID) ([]*accounting_models.PurchaseInvoiceLine, error) {
	var lines []*accounting_models.PurchaseInvoiceLine

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We filter by purchase_invoice_id and order by created_at or sequence
	err := db.WithContext(ctx).
		Where("purchase_invoice_id = ?", purchaseInvoiceID).
		Order("created_at ASC").
		Find(&lines).Error

	if err != nil {
		return nil, err
	}

	return lines, nil
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                       Purchase Payments Repo Methods                       */
/* -------------------------------------------------------------------------- */

func (r *accountingRepository) CreatePurchasePayment(ctx context.Context, payment *accounting_models.PurchasePayment) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Using WithContext to ensure the DB respects request timeouts
	if err := db.WithContext(ctx).Create(payment).Error; err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetPurchasePaymentByID(ctx context.Context, id uuid.UUID) (*accounting_models.PurchasePayment, error) {
	var payment accounting_models.PurchasePayment

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) ListPurchasePaymentsByInvoice(ctx context.Context, purchaseInvoiceID uuid.UUID) ([]*accounting_models.PurchasePayment, error) {
	var payments []*accounting_models.PurchasePayment

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We order by PaymentDate to show the timeline of payments
	err := db.WithContext(ctx).
		Where("purchase_invoice_id = ?", purchaseInvoiceID).
		Order("payment_date ASC, created_at ASC").
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) ListPurchasePaymentsBySupplier(ctx context.Context, supplierID uuid.UUID) ([]*accounting_models.PurchasePayment, error) {
	var payments []*accounting_models.PurchasePayment

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Retrieve all payments associated with the specific supplier
	// Ordered by payment date descending to show most recent activity first
	err := db.WithContext(ctx).
		Where("supplier_id = ?", supplierID).
		Order("payment_date DESC, created_at DESC").
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

// ── Aggregates / Reporting ────────────────────────────────────────────────

func (r *accountingRepository) GetSupplierOutstandingBalance(ctx context.Context, supplierID uuid.UUID, accountingBookID uuid.UUID) (float64, error) {
	var totalOutstanding float64

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We sum the amount_due for all invoices that aren't fully paid
	err := db.WithContext(ctx).
		Model(&accounting_models.PurchaseInvoice{}).
		Where("supplier_id = ? AND accounting_book_id = ? AND payment_status != 'PAID'", supplierID, accountingBookID).
		Select("COALESCE(SUM(amount_due), 0)").
		Scan(&totalOutstanding).Error

	return totalOutstanding, err
}

func (r *accountingRepository) GetPurchaseInvoiceSummary(ctx context.Context, accountingBookID uuid.UUID, fiscalYearID uuid.UUID) (*PurchaseInvoiceSummary, error) {
	var summary PurchaseInvoiceSummary
	now := time.Now()

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Base query for the specific book and year
	baseQuery := db.WithContext(ctx).
		Model(&accounting_models.PurchaseInvoice{}).
		Where("accounting_book_id = ? AND fiscal_year_id = ?", accountingBookID, fiscalYearID)

	// Execute aggregate calculations
	err := baseQuery.Select(`
        COUNT(*) as invoice_count,
        COALESCE(SUM(total_amount), 0) as total_invoiced,
        COALESCE(SUM(amount_paid), 0) as total_paid,
        COALESCE(SUM(amount_due), 0) as total_outstanding,
        COALESCE(SUM(CASE WHEN due_date < ? AND payment_status != 'PAID' THEN amount_due ELSE 0 END), 0) as total_overdue
    `, now).Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *accountingRepository) ListOverduePurchaseInvoices(ctx context.Context, accountingBookID uuid.UUID, asOf time.Time) ([]*accounting_models.PurchaseInvoice, error) {
	var invoices []*accounting_models.PurchaseInvoice

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("accounting_book_id = ?", accountingBookID).
		Where("due_date < ? AND payment_status != 'PAID'", asOf).
		Order("due_date ASC").
		Find(&invoices).Error

	if err != nil {
		return nil, err
	}

	return invoices, nil
}
