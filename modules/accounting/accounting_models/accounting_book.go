package accounting_models

import (
	"sharpnr-engine/core/models"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

type AccountingBook struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrgID        uuid.UUID `gorm:"type:uuid;not null;index"`
	Organization models.Organization

	Name string `gorm:"not null"`

	Department *string // optional can be null
	Module     *string // optional can be null

	FiscalYears   []FiscalYear   `gorm:"foreignKey:AccountingBookID"`
	AccountGroups []AccountGroup `gorm:"foreignKey:AccountingBookID"`

	Status accountingconstants.AccountingBookStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}
