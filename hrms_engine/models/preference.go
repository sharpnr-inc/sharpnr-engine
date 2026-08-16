package hrmsEngineModels

import (
	hrmsEngineConstant "sharpnr-engine/hrms_engine/constant"
	"time"

	"github.com/google/uuid"
)

// OrgHRMSPreference is one organization's HRMS configuration — a single row
// per org, created with defaults the first time the org touches HR.
//
// Deliberately NOT a schema switch. Every HR table carries a cost center
// regardless of HRModel, because every employee already has a posting in
// employee.employee_cost_centers, so the column is always populatable — even
// for a single-office org. A centralized HR simply holds a type-level policy,
// the scope resolver returns "all", and no filter is applied. That keeps one
// schema and one code path, and makes growing from CENTRALIZED to DISTRIBUTED
// a change of policy attachments rather than a backfill across every HR table.
//
// Kept small on purpose. It will accrue flags as attendance, payroll and
// recruitment land; each should earn its place the way these two do.
type OrgHRMSPreference struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// One row per organization.
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	HRModel hrmsEngineConstant.HRModel `gorm:"type:varchar(20);default:'CENTRALIZED';not null"`

	// RequireCentralApprovalForOnboarding only bites in DISTRIBUTED mode: may a
	// local HR complete an onboarding on their own, or does central HR sign off?
	// Ignored under CENTRALIZED, where there is no local HR to gate.
	RequireCentralApprovalForOnboarding bool `gorm:"default:false;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrgHRMSPreference) TableName() string {
	return "hrms.org_hrms_preference"
}

/* ---------------------------------- <***> --------------------------------- */
