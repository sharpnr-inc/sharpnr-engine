package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookID     uuid.UUID `gorm:"type:uuid;index;not null"`
	Name       string    `gorm:"not null"`
	Code       string    `gorm:"uniqueIndex;not null"`
	GroupID    uuid.UUID `gorm:"type:uuid;index"`
	Type       string
	OpeningBal float64
	CurrentBal float64
	IsDefault  bool
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
