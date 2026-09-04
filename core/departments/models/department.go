package orgDepartmentModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// this is the main table for the Departments
// a organization can have multiple departments

// OrgDeparment is an ORG-LEVEL department: it hangs off the organization
// directly, which is the track used by organizations that do not manage cost
// centers. An org that does manage them defines its departments in
// core.cost_center_departments instead; the two tracks never mix.
type OrgDeparment struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name           string    `gorm:"type:varchar(150);not null"`
	Description    string    `gorm:"type:text"`

	// Code is the short human key ("FIN"), derived from the name when the
	// caller does not supply one. Unique per organization among live rows.
	Code string `gorm:"type:varchar(50);not null;index"`

	// HeadEmployeeID is the employee who runs the department. Nullable: a
	// department is usually created before anyone is put in charge of it.
	HeadEmployeeID *uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (OrgDeparment) TableName() string {
	return "core.org_departments"
}

// OrgDepartmentEmployee is the membership row: which employees belong to an
// org-level department. OrganizationID is denormalised from the department so
// every tenant-scoped read here is a single indexed predicate.
type OrgDepartmentEmployee struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID  uuid.UUID `gorm:"type:uuid;not null;index"`
	OrgDepartmentID uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeeID      uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (OrgDepartmentEmployee) TableName() string {
	return "core.org_departments_employees"
}
