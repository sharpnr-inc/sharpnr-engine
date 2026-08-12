package orgCostCenterModels

import (
	orgCostCenterConstants "sharpnr-engine/core/cost_center/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CostCenterDepartment is a department belonging to one cost center — the
// Bangalore office's Finance team. A cost center has as many as it needs.
//
// These are the cost center's OWN departments, not placements of an org-level
// department. The two department tracks are independent by design: an org that
// manages cost centers defines its departments here, and an org that does not
// uses core/departments, where departments hang off the organization directly.
type CostCenterDepartment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// OrganizationID is denormalised from the cost center so every tenant-scoped
	// read of this table is a single indexed predicate rather than a join.
	OrganizationID  uuid.UUID `gorm:"type:uuid;not null;index"`
	OrgCostCenterID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Code is the short human key ("FIN"). Unique per cost center — enforce
	// with a composite unique index on (org_cost_center_id, code).
	Code        string `gorm:"type:varchar(50);not null;index"`
	Name        string `gorm:"type:varchar(150);not null"`
	Description string `gorm:"type:text"`

	// HeadEmployeeID is the employee who runs this department at this location.
	HeadEmployeeID *uuid.UUID `gorm:"type:uuid;index"`

	Status orgCostCenterConstants.CostCenterStatus `gorm:"type:varchar(20);default:'ACTIVE';not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CostCenterDepartment) TableName() string {
	return "core.cost_center_departments"
}

/* ---------------------------------- <***> --------------------------------- */
