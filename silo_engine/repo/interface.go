package siloEngineRepo

import (
	"context"
	siloEngineModels "sharpnr-engine/silo_engine/models"

	"github.com/google/uuid"
)

type SiloEngineRepoInterface interface {
	// InitializeSiloAccount starts the service for an org: it writes the one
	// account row every other Silo row hangs off.
	InitializeSiloAccount(ctx context.Context, acc siloEngineModels.OrgSiloAccount) error
	CreateSettingForSiloAccount(ctx context.Context, accSetting siloEngineModels.SiloAccountSetting) error
	GetSiloAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*siloEngineModels.OrgSiloAccount, error)

	// SiloAccountSwitch is the kill switch, and the path a start on an already
	// provisioned-but-disabled org takes: the account is re-activated, never
	// replaced.
	SiloAccountSwitch(ctx context.Context, orgID, orgSiloAccountID uuid.UUID, isActive bool) error
}
