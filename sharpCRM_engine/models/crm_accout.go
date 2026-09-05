package sharpCrmEngineModels

import (
	"time"

	"github.com/google/uuid"
)

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

type SharpCrmAccountSetting struct {
	ID                   uuid.UUID
	OrgSharpCrmAccountID uuid.UUID
	OrganizationID       uuid.UUID

	// we will add as we move forward
	// as we need

	LastUpdatedAt time.Time
	
}

func (SharpCrmAccountSetting) TableName() string {
	return "sharpcrm_engine.sharpcrm_account_settings"
}
