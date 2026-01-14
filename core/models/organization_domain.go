package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationDomain struct {
	ID             uuid.UUID `gorm:"not null;primaryKey"`
	OrganizationID uuid.UUID
	Domain         string // Hospitality // Medical // Retails
	Status         string // ENABLED, DISABLED
	CreatedAt      time.Time
}
