package accounting_repo

import (
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	"time"

	"github.com/google/uuid"
)

func (r *accountingRepository) GetLedger(
	accountID uuid.UUID,
	from, to time.Time,
) ([]accountingdtos.LedgerLineDto, error) {

	var lines []accountingdtos.LedgerLineDto

	err := r.db.
		Table("accounting_engine.journal_line AS jl").
		Select(`
			je.entry_date,
			je.reference,
			je.narration,
			je.entry_type,
			jl.debit,
			jl.credit,
			SUM(jl.debit - jl.credit)
			OVER (
				ORDER BY je.entry_date, je.created_at
				ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
			) AS running_balance
		`).
		Joins(`
			JOIN accounting_engine.journal_entry AS je
			ON je.id = jl.journal_entry_id
		`).
		Where("jl.account_id = ?", accountID).
		Where("je.entry_date BETWEEN ? AND ?", from, to).
		Order("je.entry_date ASC, je.created_at ASC").
		Scan(&lines).Error

	return lines, err
}
