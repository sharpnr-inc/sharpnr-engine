package orgCostCenterModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgCostCenter struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	City           string
	State          string
	Country        string

	Code string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
