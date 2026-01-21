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

	AccountingBook AccountingBook `gorm:"foreignKey:AccountingBookID"`

	Accounts []Account `gorm:"foreignKey:AccountGroupID"`

	Type accountingconstants.AccountGroupType `gorm:"type:varchar(20);not null"`

	CreatedAt time.Time

	UpdatedAt time.Time
}
