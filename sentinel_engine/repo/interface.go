package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// SentinelRepository is the data-access layer for the authorization tables.
// Each method participates in an ambient transaction when one is present in the
// context (see dbtx), so a service can run several writes atomically.
type SentinelRepositoryInterface interface {
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error
	GetRoleByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error) // this will return the Role Details
	// AssignUserRole inserts the (org, user, role) membership edge. orgID is the
	// tenant scope and must equal the role's OrganizationID.
	AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error

	// this will give us the roleID and then We can Get the Role Details based on roleID
	GetUserRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelengine_models.UserRole, error) // this is will return the role details of the employee

	// --- Managed policies (authoring layer) + generator support ---

	CreateManagedPolicy(ctx context.Context, mp *sentinelengine_models.ManagedPolicy) error
	AttachRolePolicy(ctx context.Context, edge *sentinelengine_models.RolePolicy) error
	DetachRolePolicy(ctx context.Context, orgID, roleID, managedPolicyID uuid.UUID) error
	ListManagedPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error)
	GetManagedPolicy(ctx context.Context, orgID, id uuid.UUID) (*sentinelengine_models.ManagedPolicy, error)
	ListRoles(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Role, error)
	ListAttachedManagedPolicies(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error)
	// DeleteGeneratedPoliciesByRole removes only the generated rows the generator
	// owns for a role (tagged with source_role_id); hand-authored policies survive.
	DeleteGeneratedPoliciesByRole(ctx context.Context, orgID, roleID uuid.UUID) error
}
