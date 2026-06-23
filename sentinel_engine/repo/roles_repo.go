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

// GetRoleByID loads a role by its primary key. Used to resolve a role's slug
// (e.g. when generating its Cedar policies) and to show role details in the UI.
func (r *sentinelRepo) GetRoleByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error) {
	var role sentinelengine_models.Role
	err := r.conn(ctx).Where("id = ?", id).First(&role).Error
	return role, err
}

// GetUserRole returns the role-assignment edge for an employee within an org.
// The caller can then load the role's details via GetRoleByID. Assumes a single
// role per user for now; switch to a list method when multi-role lands.
func (r *sentinelRepo) GetUserRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelengine_models.UserRole, error) {
	var ur sentinelengine_models.UserRole
	err := r.conn(ctx).
		Where("organization_id = ? AND user_id = ?", orgID, emplID).
		First(&ur).Error
	return ur, err
}
