package accounting_models

import (
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

type AccountingSetting struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID        uuid.UUID
	Organization models.Organization

	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Currency         string
	Rounding         string
	DateFormat       string
}
