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
}
