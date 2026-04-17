package accounting_engine_helpers

import (
	"errors"
	"math"

	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
)

const epsilon = 0.000001

func ValidateJournalBalance(lines []accountingdtos.JournalLineExternalRequestDto) error {

	if len(lines) == 0 {
		return errors.New("journal entry must contain at least one line")
	}

	var debitTotal float64
	var creditTotal float64

	for _, l := range lines {

		if l.Debit < 0 || l.Credit < 0 {
			return errors.New("debit or credit cannot be negative")
		}

		if (l.Debit == 0 && l.Credit == 0) ||
			(l.Debit > 0 && l.Credit > 0) {
			return errors.New("each journal line must contain either debit or credit")
		}

		debitTotal += l.Debit
		creditTotal += l.Credit
	}

	if math.Abs(debitTotal-creditTotal) > epsilon {
		return errors.New("journal entry is not balanced")
	}

	return nil
}
