package hrmsEngineRepo

import (
	"context"
	hrmsEngineModels "sharpnr-engine/hrms_engine/models"

	"github.com/google/uuid"
)

type HrmsEngineRepoInterface interface {
	// InitializeHrmsAccount starts the service for an org: it writes the one
	// account row every other HRMS row hangs off.
	InitializeHrmsAccount(ctx context.Context, acc hrmsEngineModels.OrgHrmsAccount) error
	CreateSettingForHrmsAccount(ctx context.Context, accSetting hrmsEngineModels.HrmsAccountSetting) error
	GetHrmsAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*hrmsEngineModels.OrgHrmsAccount, error)

	// HrmsAccountSwitch is the kill switch, and the path a start on an already
	// provisioned-but-disabled org takes: the account is re-activated, never
	// replaced.
	HrmsAccountSwitch(ctx context.Context, orgID, orgHrmsAccountID uuid.UUID, isActive bool) error
}
