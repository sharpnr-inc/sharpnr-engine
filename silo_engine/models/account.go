package siloEngineModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgSiloAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	IsActive       bool
	CreatedAt      time.Time
}

func (OrgSiloAccount) TableName() string {
	return "silo_engine.org_silo_accounts"
}

// SiloAccountSetting holds the account's tunables — regions, defaults,
// retention. It is provisioned in the same transaction as the account, so every
// setting added later is an ALTER plus an UPDATE and never a backfill of
// missing rows.
type SiloAccountSetting struct {
	ID               uuid.UUID
	OrgSiloAccountID uuid.UUID
	OrganizationID   uuid.UUID

	// Tagged because the field name is misspelled: the default naming strategy
	// would look for an updted_at column, and the schema is not going to carry
	// the typo. GORM only auto-manages a field literally named UpdatedAt, so
	// this one is set explicitly on write.
	UpdtedAt time.Time `gorm:"column:updated_at"`
	// we will add as we go further
}

func (SiloAccountSetting) TableName() string {
	return "silo_engine.silo_account_settings"
}
