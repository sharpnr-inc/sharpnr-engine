package folioEngineRepo

import (
	"context"
	folioEngineModels "sharpnr-engine/folio_engine/models"

	"github.com/google/uuid"
)

type FolioEngineRepoInterface interface {
	// InitializeFolioAccount starts the service for an org: it writes the one
	// account row every other Folio row hangs off.
	InitializeFolioAccount(ctx context.Context, acc folioEngineModels.OrgFolioAccount) error
	CreateSettingForFolioAccount(ctx context.Context, accSetting folioEngineModels.FolioAccountSetting) error
	GetFolioAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*folioEngineModels.OrgFolioAccount, error)

	// FolioAccountSwitch is the kill switch, and the path a start on an already
	// provisioned-but-disabled org takes: the account is re-activated, never
	// replaced.
	FolioAccountSwitch(ctx context.Context, orgID, orgFolioAccountID uuid.UUID, isActive bool) error
}
