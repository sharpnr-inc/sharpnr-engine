package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

type AccountingBook struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// This is the Main Organization ID
	// or you can say its tenantID
	OrgID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name string `gorm:"not null"`

	Department *string // optional can be null
	Module     *string // optional can be null

	FiscalYears   []FiscalYear   `gorm:"foreignKey:AccountingBookID"`
	AccountGroups []AccountGroup `gorm:"foreignKey:AccountingBookID"`

	Status accountingconstants.AccountingBookStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}
