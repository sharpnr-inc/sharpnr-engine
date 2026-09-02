package orgCostCenterModels

import (
	"time"

	"github.com/google/uuid"
)

// A Cost Center Can multiple Employees

type OrgCostCenterEmployee struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgCostCenterID uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeeID      uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrgCostCenterEmployee) TableName() string {
	return "core.org_cost_center_employees"
}
