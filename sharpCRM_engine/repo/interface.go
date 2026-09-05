package sharpCrmEngineRepo

import (
	"context"
	sharpCrmEngineModels "sharpnr-engine/sharpCRM_engine/models"

	"github.com/google/uuid"
)

type SharpCrmEngineRepoInterace interface {
	// this method will start the service or create the account table
	InitializeCrmAccount(ctx context.Context, acc sharpCrmEngineModels.OrgSharpCrmAccount) error
	CreateSettingForCrmAccount(ctx context.Context, acc_setting sharpCrmEngineModels.SharpCrmAccountSetting) error
	GetCrmAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*sharpCrmEngineModels.OrgSharpCrmAccount, error)

	// CrmAccountSwitch is the kill switch, and the path a start on an already
	// provisioned-but-disabled org takes: the account is re-activated, never
	// replaced.
	CrmAccountSwitch(ctx context.Context, orgID, orgCrmAccountID uuid.UUID, isActive bool) error
}
