package orgDepartmentModels

import (
	"time"

	"github.com/google/uuid"
)

// this is the main table for the Departments
// a organization can have multiple departments

type OrgDeparment struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Description    string
	Code           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (OrgDeparment) TableName() string {
	return "core.departments"
}


