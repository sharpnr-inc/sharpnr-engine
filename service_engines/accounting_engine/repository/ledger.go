package accounting_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	"time"

	"github.com/google/uuid"
)

func (r *accountingRepository) GetLedgerEntries(
	ctx context.Context,
	bookID uuid.UUID,
	fiscalYearID uuid.UUID,
	accountID *uuid.UUID,
	page int,
	limit int,
) ([]accountingdtos.LedgerEntryDTO, int64, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	offset := (page - 1) * limit

	query := db.Table("accounting_engine.journal_line jl").
		Joins("JOIN accounting_engine.journal_entry je ON je.id = jl.journal_entry_id").
		Where("je.accounting_book_id = ?", bookID).
		Where("je.fiscal_year_id = ?", fiscalYearID)

	// optional account filter
	if accountID != nil {
		query = query.Where("jl.account_id = ?", *accountID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var ledger []accountingdtos.LedgerEntryDTO

	err := query.
		Select(`
			je.id,
			je.entry_date,
			je.reference,
			je.narration,
			jl.debit,
			jl.credit
		`).
		Order("je.entry_date, je.created_at").
		Limit(limit).
		Offset(offset).
		Scan(&ledger).Error

	if err != nil {
		return nil, 0, err
	}

	return ledger, total, nil
}

func (r *accountingRepository) GetLedger(
	accountID uuid.UUID,
	from, to time.Time,
) ([]accountingdtos.LedgerLineDto, error) {

	var lines []accountingdtos.LedgerLineDto

	r.db.Raw(`
        SELECT
            je.entry_date,
            je.reference,
            je.narration,
            je.entry_type,
            jl.debit,
            jl.credit,
            SUM(jl.debit - jl.credit)
                OVER (ORDER BY je.entry_date, je.created_at
                      ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
                ) AS running_balance
        FROM accounting_engine.journal_line  jl
        JOIN accounting_engine.journal_entry je
          ON je.id = jl.journal_entry_id
        WHERE jl.account_id = ?
          AND je.entry_date BETWEEN ? AND ?
        ORDER BY je.entry_date ASC, je.created_at ASC
    `, accountID, from, to).Scan(&lines)

	return lines, nil
}
