package accounting_models

import "github.com/google/uuid"

type AccountingSetting struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID      uuid.UUID
	Currency   string
	Rounding   string
	DateFormat string
}
