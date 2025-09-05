package models

import (
	"sharpnr-engine/core/constants"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID             uuid.UUID
	Name           constants.EmployeeRoles
	Description    string
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
