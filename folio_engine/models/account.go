package folioEngineModels

import (
	"time"

	"github.com/google/uuid"
)

// An organization only gonna have one Folio Account.
// The reason for this Account table is because one day we might need an
// individual (non-org) Folio account.
//
// IsActive is the kill switch: switching the service off leaves every document
// in place, and starting it again flips this back rather than minting a second
// account.
type OrgFolioAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgFolioAccount) TableName() string {
	return "folio_engine.org_folio_accounts"
}

// FolioAccountSetting holds the account's tunables — default sharing,
// templates, retention. It is provisioned in the same transaction as the
// account, so every setting added later is an ALTER plus an UPDATE and never a
// backfill of missing rows.
type FolioAccountSetting struct {
	ID                uuid.UUID
	OrgFolioAccountID uuid.UUID
	OrganizationID    uuid.UUID

	// we will add as we move forward

	UpdatedAt time.Time
}

func (FolioAccountSetting) TableName() string {
	return "folio_engine.folio_account_settings"
}
