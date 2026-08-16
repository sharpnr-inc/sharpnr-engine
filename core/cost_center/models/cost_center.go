package orgCostCenterModels

import (
	orgCostCenterConstants "sharpnr-engine/core/cost_center/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrgCostCenter is one operating location of an organization — the Bangalore
// office, the Dubai branch, a warehouse. An organization has as many as it
// needs, in as many countries as it needs.
//
// A cost center owns its own departments (see CostCenterDepartment). An org
// that does not want to manage cost centers at all uses core/departments
// instead, where departments hang directly off the organization.
type OrgCostCenter struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// LegalEntityID is the registered entity this cost center operates under.
	// Nullable because a cost center can be stood up before its entity is
	// registered, but it must be set before anything here can be invoiced or
	// filed — an office in another country almost always belongs to a different
	// entity than head office, and that is what decides tax treatment.
	LegalEntityID *uuid.UUID `gorm:"type:uuid;index"`

	// Code is the short human key ("BLR-01"). Unique per organization —
	// enforce with a composite unique index in the migration, not here.
	Code        string `gorm:"type:varchar(50);not null;index"`
	Name        string `gorm:"type:varchar(150);not null"`
	Description string `gorm:"type:text"`

	Type   orgCostCenterConstants.CostCenterType   `gorm:"type:varchar(30);index"`
	Status orgCostCenterConstants.CostCenterStatus `gorm:"type:varchar(20);default:'ACTIVE';not null;index"`

	// Address
	AddressLine string `gorm:"type:text"`
	City        string `gorm:"type:varchar(100);index"`
	State       string `gorm:"type:varchar(100);index"`
	Country     string `gorm:"type:varchar(100);index"`
	Pincode     string `gorm:"type:varchar(20);index"`

	// Timezone of the site, e.g. "Asia/Dubai". Attendance cutoffs, shift
	// boundaries and payroll periods are all local to the office, so this
	// cannot be inferred from the organization.
	Timezone string `gorm:"type:varchar(50)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (OrgCostCenter) TableName() string {
	return "core.org_cost_centers"
}

/* ---------------------------------- <***> --------------------------------- */
