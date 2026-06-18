package sentinelservice

import (
	"context"

	"github.com/google/uuid"
)

// SentinelService is the authorization service injected into other modules.
// For now it exposes provisioning (used at signup). Evaluation and grid
// management will be added here as those features land.
type SentinelService interface {
	// ProvisionSuperAdmin creates the org's system "Super Admin" role with full
	// access and assigns it to the given employee. Returns the new role's ID.
	ProvisionSuperAdmin(ctx context.Context, orgID, employeeID uuid.UUID) (uuid.UUID, error)
}
