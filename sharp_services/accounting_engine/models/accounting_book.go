package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type AccountingBook struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID            uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationLegalEntityID uuid.UUID

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

func (AccountingBook) TableName() string {
	return "accounting_engine.accounting_book"
}
