package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

type AccountingBook struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrgID          uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID uuid.UUID

	Name string `gorm:"not null"`

	Domain string

	Module *string

	SubModule *string

	FiscalYears []FiscalYear `gorm:"foreignKey:AccountingBookID"`

	AccountGroups []AccountGroup `gorm:"foreignKey:AccountingBookID"`

	Status accountingconstants.AccountingBookStatus

	CreatedAt time.Time

	UpdatedAt time.Time
}
