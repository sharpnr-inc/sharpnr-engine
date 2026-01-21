package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type FiscalYear struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrgID uuid.UUID

	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	AccountingBook   AccountingBook

	Name string

	StartDate time.Time
	EndDate   time.Time

	Status    accountingconstants.FiscalYearType
	CreatedAt time.Time

	UpdatedAt time.Time
}
