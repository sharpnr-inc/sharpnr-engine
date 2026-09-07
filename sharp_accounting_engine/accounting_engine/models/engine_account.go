package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

// An organization only gonna have one accounting-engine account. It is the
// switch that says the org has taken the service — distinct from the
// AccountingBook rows, which hang off each legal entity.
//
// IsActive is the kill switch: switching the service off leaves every ledger
// row in place, and starting it again flips this back rather than minting a
// second account.
type OrgAccountineEngineAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgAccountineEngineAccount) TableName() string {
	return "accounting_engine.org_accounting_engine_accounts"
}

// AccountingEngineAccountSetting holds the account's org-wide tunables. It is
// provisioned in the same transaction as the account, so every setting added
// later is an ALTER plus an UPDATE and never a backfill of missing rows.
//
// Distinct from AccountingBookSetting, which configures one legal entity's
// book; this row configures the service for the whole organization.
type AccountingEngineAccountSetting struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	// Tagged because the field name carries a typo ("Accountine"): the default
	// naming strategy would look for org_accountine_engine_account_id, and the
	// schema is not going to carry it.
	OrgAccountineEngineAccountID uuid.UUID `gorm:"column:org_accounting_engine_account_id"`

	UpdatedAt time.Time

	// we will add as we move forward
}

func (AccountingEngineAccountSetting) TableName() string {
	return "accounting_engine.accounting_engine_account_settings"
}
