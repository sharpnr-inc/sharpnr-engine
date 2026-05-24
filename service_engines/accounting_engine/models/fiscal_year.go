package accounting_models

import (
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type FiscalYear struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name             string
	StartDate        time.Time
	EndDate          time.Time
	Status           accountingconstants.FiscalYearType
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (FiscalYear) TableName() string {
	return "accounting_engine.fiscal_year"
}
