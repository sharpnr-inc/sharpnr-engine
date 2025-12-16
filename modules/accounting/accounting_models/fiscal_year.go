package accounting_models

import (
	"sharpnr-engine/core/models"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

type FiscalYear struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrgID        uuid.UUID
	Organization models.Organization

	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	AccountingBook   AccountingBook

	Name string

	StartDate time.Time
	EndDate   time.Time

	Status    accountingconstants.FiscalYearType
	CreatedAt time.Time

	UpdatedAt time.Time
}
