package siloEngineModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgSiloAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgSiloAccount) TableName() string {
	return "silo_engine.org_silo_accounts"
}

type SiloAccountSetting struct {
	ID               uuid.UUID
	OrgSiloAccountID uuid.UUID
	OrganizationID   uuid.UUID
	UpdtedAt         time.Time
	// we will add as we go further
}

func (SiloAccountSetting) TableName() string {
	return "silo_engine.silo_account_settings"
}
