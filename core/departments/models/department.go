package orgDepartmentModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (OrgDeparment) TableName() string {
	return "core.org_departments"
}

type OrgDepartmentEmployee struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	OrgDepartmentID uuid.UUID
	EmployeeID      uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (OrgDepartmentEmployee) TableName() string {
	return "core.org_departments_employees"
}
