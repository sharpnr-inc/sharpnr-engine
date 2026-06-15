package accounting_engine_helpers

import (
	"fmt"
	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"

	"time"

	"github.com/google/uuid"
)

// GenerateIndianFiscalYear creates a FiscalYear model based on Indian standards
func GenerateIndianFiscalYear(bookID uuid.UUID) *accounting_models.FiscalYear {

	now := time.Now()
	var startYear int

	// Indian FY: April to March
	if now.Month() >= time.April {
		startYear = now.Year()
	} else {
		startYear = now.Year() - 1
	}

	endYear := startYear + 1

	// Set Start: April 1st, 00:00:00
	startDate := time.Date(startYear, time.April, 1, 0, 0, 0, 0, time.Local)
	// Set End: March 31st, 23:59:59 (next year)
	endDate := time.Date(endYear, time.March, 31, 23, 59, 59, 0, time.Local)

	return &accounting_models.FiscalYear{
		ID:               uuid.New(),
		AccountingBookID: bookID,
		Name:             fmt.Sprintf("FY %d-%02d", startYear, endYear%100),
		StartDate:        startDate,
		EndDate:          endDate,
		Status:           accountingconstants.Open,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
