package accounting_repo

import (
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	"time"

	"github.com/google/uuid"
)

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
