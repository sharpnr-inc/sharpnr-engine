package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type AccountingBook struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID        uuid.UUID `gorm:"type:uuid;not null;index"`
	Name         string    `gorm:"not null"`
	Department   string
	FiscalYearID uuid.UUID `gorm:"type:uuid"`
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
