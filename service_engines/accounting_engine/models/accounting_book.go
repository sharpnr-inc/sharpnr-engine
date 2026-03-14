package accounting_models

import (
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

// This is whole wrapper or top level for Accounting Book
// Accounting book have multiple Accounting Group
// have multiple Fiscal Year.
// Accounting group is the Collection of Accounts

type AccountingBook struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID            uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name string `gorm:"not null"`

	Status accountingconstants.AccountingBookStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AccountingBook) TableName() string {
	return "accounting_engine.accounting_book"
}

type AccountingBookMeta struct {
	ID               uuid.UUID
	AccountingBookID uuid.UUID
	Currency         string
}
