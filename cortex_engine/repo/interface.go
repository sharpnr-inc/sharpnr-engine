package cortexEngineRepo

import (
	"context"
	cortexEngineModels "sharpnr-engine/cortex_engine/models"

	"github.com/google/uuid"
)

type CortexEngineRepoInterface interface {
	// InitializeCortexAccount starts the service for an org: it writes the one
	// account row every other Cortex row hangs off.
	InitializeCortexAccount(ctx context.Context, acc cortexEngineModels.OrgCortexAccount) error
	CreateSettingForCortexAccount(ctx context.Context, accSetting cortexEngineModels.CortexAccountSetting) error
	GetCortexAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*cortexEngineModels.OrgCortexAccount, error)

	// CortexAccountSwitch is the kill switch, and the path a start on an already
	// provisioned-but-disabled org takes: the account is re-activated, never
	// replaced.
	CortexAccountSwitch(ctx context.Context, orgID, orgCortexAccountID uuid.UUID, isActive bool) error
}
