package cablesEngineRepo

import (
	"context"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
)

// CablesRepoInterface is the cables module's persistence port.
//
// Every method carries organizationID because the tenant predicate belongs on
// the query, not on the caller's memory: a channel id from another org must
// read as "not found" rather than resolve.
type CablesRepoInterface interface {
	// IntializeOrgCablesAccount mints the org's one cables account and returns
	// it. The unique index on organization_id is what makes a concurrent second
	// call fail rather than produce a duplicate.
	IntializeOrgCablesAccount(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error)

	// GetCablesAccountDetails returns the org's account, or
	// ErrCablesServiceNotInitialized when the org has never started cables.
	GetCablesAccountDetails(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error)

	// CablesAccountSwitch turns the org's cables service on or off. isActive is
	// an explicit target rather than a toggle so a retried request lands on the
	// state the caller asked for instead of flipping it back.
	CablesAccountSwitch(ctx context.Context, orgID, orgCablesAccountID uuid.UUID, isActive bool) error

	/* -------------------------------- channels ------------------------------- */

	CreateChannel(ctx context.Context, channel cables_engine_models.CablesChannel) error

	// ListChannels returns the account's channels in sidebar order.
	ListChannels(ctx context.Context, orgID, orgCablesAccountID uuid.UUID) ([]cables_engine_models.CablesChannel, error)

	GetChannelByID(ctx context.Context, orgID, channelID uuid.UUID) (*cables_engine_models.CablesChannel, error)

	// ChannelSlugExists reports whether a live channel in the account already
	// holds this slug. Soft-deleted rows do not count, matching the partial
	// unique index that enforces the same rule underneath.
	ChannelSlugExists(ctx context.Context, orgCablesAccountID uuid.UUID, slug string) (bool, error)
}
