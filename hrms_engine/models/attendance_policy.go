package hrmsEngineModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendancePolicy is the set of rules used to turn raw punches into a judged
// day — present, half day, late, absent, overtime. It says HOW work is scored;
// Shift says when work is expected.
//
// Scope is the resolution rule the whole HRMS uses: a row with a nil
// OrgCostCenterID is the organization default, and a row naming a cost center
// overrides it for that site. A CENTRALIZED org has exactly one row with a nil
// scope; a DISTRIBUTED org adds per-site rows. Nothing about the schema changes
// between the two — the difference is which rows exist and who Sentinel lets
// create them.
//
// Central HR governs local HR's ability to add site overrides through Sentinel
// policy (a type-level grant vs a per-cost-center one), NOT through a column
// here. Permissions are not business rules and must not be modelled twice.
type AttendancePolicy struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// nil = organization default. Set = override for that cost center.
	OrgCostCenterID *uuid.UUID `gorm:"type:uuid;index"`

	Code string `gorm:"type:varchar(50);not null;index"`
	Name string `gorm:"type:varchar(150);not null"`

	/* ------------------------------- Lateness ------------------------------- */

	// GracePeriodMinutes is how late an arrival may be before it counts at all.
	GracePeriodMinutes int `gorm:"type:int;default:0;not null"`

	// MaxLateMarksPerMonth is how many late arrivals are tolerated before the
	// policy converts them into a deduction. Zero means late marks are recorded
	// but never converted.
	MaxLateMarksPerMonth int `gorm:"type:int;default:0;not null"`

	/* ------------------------------ Day scoring ----------------------------- */

	// MinMinutesForHalfDay / MinMinutesForFullDay are measured on time actually
	// worked, break excluded. Below the half-day figure the day is absent.
	MinMinutesForHalfDay int `gorm:"type:int;not null"`
	MinMinutesForFullDay int `gorm:"type:int;not null"`

	/* ------------------------------- Overtime ------------------------------- */

	OvertimeEnabled bool `gorm:"default:false;not null"`

	// OvertimeAfterMinutes is measured beyond the shift's expected duration, so
	// it does not need restating per shift.
	OvertimeAfterMinutes int `gorm:"type:int;default:0;not null"`

	/* ---------------------------- Regularisation ---------------------------- */

	// Whether an employee may ask HR to correct a missed or wrong punch, and how
	// often. This is the request's existence, not who may approve it — approval
	// rights are Sentinel's.
	RegularisationEnabled      bool `gorm:"default:true;not null"`
	MaxRegularisationsPerMonth int  `gorm:"type:int;default:0;not null"`

	IsActive bool `gorm:"default:true;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (AttendancePolicy) TableName() string {
	return "hrms.attendance_policy"
}

/* ---------------------------------- <***> --------------------------------- */
