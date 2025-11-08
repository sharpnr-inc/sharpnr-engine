package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type FiscalYear struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID     uuid.UUID
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
