package cortexEngineModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgCortexAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	IsActive       bool
}

func (OrgCortexAccount) TableName() string {
	return "cortex_engine.org_cortex_accounts"
}

type CortexAccountSetting struct {
	ID                 uuid.UUID
	OrgCortexAccountID uuid.UUID
	OrganizationID     uuid.UUID
	UpdatedAt          time.Time

	// we will add as we move forward
}

func (CortexAccountSetting) TableName() string {
	return "cortex_engine.cortex_account_settings"
}
