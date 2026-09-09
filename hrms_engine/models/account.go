package hrmsEngineModels

import (
	"time"

	"github.com/google/uuid"
)

// An organization only gonna have one HRMS Account.
// The reason for this Account table is because one day we might need an
// individual (non-org) HRMS account.
//
// IsActive is the kill switch: switching the service off leaves every employee
// record in place, and starting it again flips this back rather than minting a
// second account.
type OrgHrmsAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgHrmsAccount) TableName() string {
	return "hrms_engine.org_hrms_accounts"
}

// HrmsAccountSetting holds the account's tunables — the attendance, payroll and
// recruitment flags as those flows land. It is provisioned in the same
// transaction as the account, so every setting added later is an ALTER plus an
// UPDATE and never a backfill of missing rows.
type HrmsAccountSetting struct {
	ID               uuid.UUID
	OrgHrmsAccountID uuid.UUID
	OrganizationID   uuid.UUID

	// we will add as we move forward

	UpdatedAt time.Time
}

func (HrmsAccountSetting) TableName() string {
	return "hrms_engine.hrms_account_settings"
}
