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
	// AssignUserRole inserts the (org, user, role) membership edge. orgID is the
	// tenant scope and must equal the role's OrganizationID.
	AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error
}
