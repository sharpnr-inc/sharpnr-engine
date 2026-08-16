package hrmsEngineModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Shift is a named work pattern — "General 09:00–18:00", "Night 22:00–06:00".
// It says WHEN work is expected. How the resulting punches are judged (grace,
// late marks, half days) is AttendancePolicy's job; the two are orthogonal so a
// cost center can run three shifts under one policy.
//
// Times are "HH:MM" strings and overnight is an explicit flag, matching the
// existing hospitality.restaurant opening/closing convention rather than
// inventing a second way to express the same thing.
type Shift struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// OrgCostCenterID nil means the shift is available org-wide; set means it
	// belongs to one site. Present from the first row so a CENTRALIZED org
	// simply leaves it nil — see OrgHRMSPreference on why scope is never
	// conditional.
	OrgCostCenterID *uuid.UUID `gorm:"type:uuid;index"`

	Code string `gorm:"type:varchar(50);not null;index"`
	Name string `gorm:"type:varchar(150);not null"`

	StartTime string `gorm:"type:varchar(5);not null"` // "09:00"
	EndTime   string `gorm:"type:varchar(5);not null"` // "18:00"

	// IsOvernight is true when EndTime falls on the following day, which is
	// what stops a 22:00–06:00 shift being read as a negative duration.
	IsOvernight bool `gorm:"default:false;not null"`

	BreakMinutes int `gorm:"type:int;default:0;not null"`

	// WorkingDaysMask is a bitmask of hrmsEngineConstant.WeekDay* values.
	WorkingDaysMask int16 `gorm:"type:smallint;not null"`

	IsActive bool `gorm:"default:true;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Shift) TableName() string {
	return "hrms.shift"
}

/* ---------------------------------- <***> --------------------------------- */

// EmployeeShiftAssignment is which shift an employee works, effective-dated the
// same way their cost center posting is: moving someone to nights closes the
// current row and opens a new one, so attendance recomputed for a past month
// still judges them against the shift they were actually on.
type EmployeeShiftAssignment struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeeID     uuid.UUID `gorm:"type:uuid;not null;index"`
	ShiftID        uuid.UUID `gorm:"type:uuid;not null;index"`

	EffectiveFrom time.Time  `gorm:"type:date;not null"`
	EffectiveTo   *time.Time `gorm:"type:date"` // nil = current assignment

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (EmployeeShiftAssignment) TableName() string {
	return "hrms.employee_shift_assignment"
}

/* ---------------------------------- <***> --------------------------------- */
