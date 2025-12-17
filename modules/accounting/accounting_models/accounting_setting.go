package accounting_models

import (
	"github.com/google/uuid"
)

type AccountingBookSetting struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID uuid.UUID

	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Currency         string
	Rounding         string
	DateFormat       string
}
