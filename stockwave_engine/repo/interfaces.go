package stockwaveEngineRepo

import (
	"context"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
)

// StockwaveEngineRepoInterface is the persistence seam for the stockwave
// engine.
//
// It carries the account layer only. Migration 000062 dropped the inventory
// tables this package used to read and write; the stock structures that replace
// them join this interface as they are designed, all hanging off the account
// row below.
type StockwaveEngineRepoInterface interface {
	// InitializeStockwaveAccount starts the service for an org: it writes the
	// one account row every other stockwave row hangs off.
	InitializeStockwaveAccount(ctx context.Context, acc stockwaveEngineModels.OrgStockwaveEngineAccounts) error
	CreateSettingForStockwaveAccount(ctx context.Context, accSetting stockwaveEngineModels.StockwaveEngineAccountSetting) error
	GetStockwaveAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*stockwaveEngineModels.OrgStockwaveEngineAccounts, error)

	// StockwaveAccountSwitch is the kill switch, and the path a start on an
	// already provisioned-but-disabled org takes: the account is re-activated,
	// never replaced.
	StockwaveAccountSwitch(ctx context.Context, orgID, orgStockwaveAccountID uuid.UUID, isActive bool) error
}
