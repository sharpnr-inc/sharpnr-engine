package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type AccountGroup struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"not null"`

	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type accountingconstants.AccountGroupType `gorm:"type:varchar(20);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AccountGroup) TableName() string {
	return "accounting_engine.account_group"
}
