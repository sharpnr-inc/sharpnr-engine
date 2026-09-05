package sharpCrmEngineRepo

import (
	"context"
	sharpCrmEngineModels "sharpnr-engine/sharpCRM_engine/models"

	"github.com/google/uuid"
)

func (d *repository) InitializeCrmAccount(ctx context.Context, acc sharpCrmEngineModels.OrgSharpCrmAccount) error {
	panic("")
}

func (d *repository) GetCrmAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*sharpCrmEngineModels.OrgSharpCrmAccount, error) {
	panic("")
}

func (d *repository) CreateSettingForCrmAccount(ctx context.Context, acc_setting sharpCrmEngineModels.SharpCrmAccountSetting) error {
	panic("")
}
