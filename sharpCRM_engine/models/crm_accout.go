package sharpCrmEngineModels

import (
	"time"

	"github.com/google/uuid"
)

// An organization only gonna have one SharpCRM Account.
// The reason for this Account table is because one day we might need an
// individual (non-org) CRM account.
//
// IsActive is the kill switch: switching the service off leaves every CRM row
// in place, and starting it again flips this back rather than minting a second
// account.
type OrgSharpCrmAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	IsActive       bool
	UpdatedAt      time.Time
}

func (OrgSharpCrmAccount) TableName() string {
	return "sharpcrm_engine.org_sharpcrm_accounts"
}

// SharpCrmAccountSetting holds the account's tunables. It is provisioned in the
// same transaction as the account, so every setting added later is an ALTER
// plus an UPDATE and never a backfill of missing rows.
type SharpCrmAccountSetting struct {
	ID uuid.UUID

	// Tagged because the default naming strategy would split "SharpCrm" into
	// two words and produce org_sharp_crm_account_id, which is not the column.
	OrgSharpCrmAccountID uuid.UUID `gorm:"column:org_sharpcrm_account_id"`
	OrganizationID       uuid.UUID

	// we will add as we move forward
	// as we need

	LastUpdatedAt time.Time
}

func (SharpCrmAccountSetting) TableName() string {
	return "sharpcrm_engine.sharpcrm_account_settings"
}
