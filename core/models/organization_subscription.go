package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationSubscription struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"index"`

	PlanCode string
	Status   string // ACTIVE, PAUSED, CANCELLED

	RenewsAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
