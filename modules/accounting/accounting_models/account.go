package accounting_models

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"not null"`
	Code string    `gorm:"uniqueIndex;not null"`

	AccountGroupID uuid.UUID `gorm:"type:uuid;index"`
	AccountGroup   AccountGroup

	Type accountingconstants.AccountType

	IsDefault bool
	Status    string

	CreatedAt time.Time
	UpdatedAt time.Time
}
