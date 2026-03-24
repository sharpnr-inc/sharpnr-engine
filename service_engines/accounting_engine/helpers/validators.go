package accounting_engine_helpers

import (
	"errors"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
)

func ValidateJournalBalance(lines []accountingdtos.JournalLineRequestDto) error {
	var debitTotal float64
	var creditTotal float64

	for _, l := range lines {
		debitTotal += l.Debit
		creditTotal += l.Credit
	}

	if debitTotal != creditTotal {
		return errors.New("journal entry is not balanced")
	}

	return nil
}
