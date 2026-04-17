package accounting_repo

import (
	"context"
	"fmt"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	"time"

	"github.com/google/uuid"
)

// repo/report.repo.go

func (r *accountingRepository) GetAccountMovementsForPeriod(
	ctx context.Context,
	bookID uuid.UUID,
	from, to time.Time,
) ([]accountingdtos.AccountMovementRow, error) {

	var rows []accountingdtos.AccountMovementRow

	err := r.db.WithContext(ctx).Raw(`
        SELECT
            a.id                AS account_id,
            a.name              AS account_name,
            a.code              AS account_code,
            a.opening_balance   AS opening_balance,
            a.opening_is_debit  AS opening_is_debit,
            ag.name       		AS group_name,
            ag.type             AS group_type,
            COALESCE(SUM(jl.debit),  0) AS total_debit,
            COALESCE(SUM(jl.credit), 0) AS total_credit
        FROM accounting_engine.account a
        JOIN accounting_engine.account_group ag
            ON ag.id = a.account_group_id
        LEFT JOIN accounting_engine.journal_line jl
            ON jl.account_id = a.id
        LEFT JOIN accounting_engine.journal_entry je
            ON je.id = jl.journal_entry_id
            AND je.accounting_book_id = ?
            AND je.entry_date BETWEEN ? AND ?
        WHERE ag.accounting_book_id = ?
        GROUP BY
            a.id, a.name, a.code, a.opening_balance,
            a.opening_is_debit, ag.name, ag.type
        ORDER BY a.code ASC
    `, bookID, from, to, bookID).
		Scan(&rows).Error

	if err != nil {
		return nil, fmt.Errorf("GetAccountMovementsForPeriod: %w", err)
	}

	return rows, nil
}
