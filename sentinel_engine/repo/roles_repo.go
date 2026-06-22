package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// CreateRole persists a role. The ID is application-generated; if the caller
// left it unset we assign one here so role.ID is valid immediately (e.g. to
// assign the role to a user in the same transaction).
func (r *sentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	if role.ID == uuid.Nil {
		role.ID = uuid.New()
	}
	return r.conn(ctx).Create(role).Error
}

// AssignUserRole inserts the (org, user, role) membership edge. It is the junction
// that makes a user a member of a role; orgID carries the tenant scope.
func (r *sentinelRepo) AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	membership := &sentinelengine_models.UserRole{
		OrganizationID: orgID,
		UserID:         userID,
		RoleID:         roleID,
	}
	return r.conn(ctx).Create(membership).Error
}
