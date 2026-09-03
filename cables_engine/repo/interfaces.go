package cablesEngineRepo

import (
	"context"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
)

type CablesRepoInterface interface {
	IntializeOrgCablesAccount(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error)
	GetCablesAccountDetails(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error) // returns the Account or err
	CablesAccountSwitch(ctx context.Context, orgID, orgCablesAccountID uuid.UUID) error
	CreateChannel(ctx context.Context, channel cables_engine_models.CablesChannel) error
}
