package stockwaveEngineModels

import (
	"time"

	"github.com/google/uuid"
)

// An organization only gonna have one stockwave-engine account. It is the
// switch that says the org has taken the service, and every stock structure the
// engine grows later hangs off it.
//
// IsActive is the kill switch: switching the service off leaves every stock row
// in place, and starting it again flips this back rather than minting a second
// account.
type OrgStockwaveEngineAccounts struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgStockwaveEngineAccounts) TableName() string {
	return "stockwave_engine.org_stockwave_engine_accounts"
}

// StockwaveEngineAccountSetting holds the account's org-wide tunables. It is
// provisioned in the same transaction as the account, so every setting added
// later is an ALTER plus an UPDATE and never a backfill of missing rows.
type StockwaveEngineAccountSetting struct {
	ID                          uuid.UUID
	OrganizationID              uuid.UUID
	OrgStockwaveEngineAccountID uuid.UUID
	UpdatedAt                   time.Time
}

func (StockwaveEngineAccountSetting) TableName() string {
	return "stockwave_engine.stockwave_engine_account_settings"
}
