package sharpCrmEngineModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgSharpCrmAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	IsActive       bool
	UpdatedAt      time.Time
}
