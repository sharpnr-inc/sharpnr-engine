package sentinelengine

import (
	"context"

	"github.com/cedar-policy/cedar-go/types"
	"github.com/google/uuid"

	sentinelengine_models "sharpnr-engine/sentinel/models"
)

// GroupMembership is one node in a user's group closure: the group, its parent
// group (empty for a top-level group), and whether the user is a *direct*
// member. Only direct memberships make the user a member; ancestor groups are
// reached transitively via the ParentSlug edge.
type GroupMembership struct {
	Slug       string
	ParentSlug string
	Direct     bool
}

// Repository is the slice of the data layer the engine needs. Your
// sentinel_repo.SentinelRepository should satisfy this directly or via a thin
// adapter. Defining it here (rather than importing the full repo interface)
// keeps the engine decoupled and easy to test with a fake.
type Repository interface {
	// ListEnabledPolicies returns every enabled policy row for an org.
	ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	// UserRoleSlugs returns the slugs of the roles the user is directly assigned.
	UserRoleSlugs(ctx context.Context, orgID, userID uuid.UUID) ([]string, error)
	// UserGroupClosure returns the user's direct groups plus all ancestor groups
	// (typically via a recursive CTE over groups.parent_group_id).
	UserGroupClosure(ctx context.Context, orgID, userID uuid.UUID) ([]GroupMembership, error)
	// UserAttributes returns the principal's ABAC attributes (e.g. department,
	// position, managerId), keyed by the name policies reference as
	// `principal.<key>`. Return an empty (non-nil) map when the user has none;
	// the engine adds `orgId` itself and it cannot be overridden here.
	UserAttributes(ctx context.Context, orgID, userID uuid.UUID) (map[string]types.Value, error)
}
