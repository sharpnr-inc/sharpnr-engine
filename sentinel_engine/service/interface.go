package sentinelservice

import (
	"context"

	"github.com/google/uuid"
)

// SentinelService is the authorization service injected into other modules.
// For now it exposes provisioning (used at signup). Evaluation and grid
// management will be added here as those features land.
type SentinelService interface {
	// ProvisionOrganization sets up a new org's authorization defaults: it creates
	// the full-access "Owner" role and its policy, then assigns the founding
	// employee to that role. Runs in one transaction (joins an ambient tx if the
	// context carries one), so an org is never left half-provisioned.
	ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error
}
