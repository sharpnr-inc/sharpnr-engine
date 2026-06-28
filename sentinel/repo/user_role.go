package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AssignRoleToUser assigns a role to a user. Idempotent: if the user already has
// the role it is a no-op.
func (r *sentinelRepo) AssignRoleToUser(ctx context.Context, organizationID, userID, roleID uuid.UUID) error {
	assigned, err := r.IsRoleAssignedToUser(ctx, userID, roleID)
	if err != nil {
		return err
	}
	if assigned {
		return nil
	}
	edge := &sentinelengine_models.UserRole{
		ID:             uuid.New(),
		OrganizationID: organizationID,
		UserID:         userID,
		RoleID:         roleID,
	}
	return r.conn(ctx).Create(edge).Error
}

// RemoveRoleFromUser removes a role assignment (hard delete — the edge has no
// soft-delete column, so the same (user, role) pair can be re-assigned later).
func (r *sentinelRepo) RemoveRoleFromUser(ctx context.Context, organizationID, userID, roleID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND user_id = ? AND role_id = ?", organizationID, userID, roleID).
		Delete(&sentinelengine_models.UserRole{}).Error
}

// GetRolesAssignedToUser returns the role-assignment edges for a user.
func (r *sentinelRepo) GetRolesAssignedToUser(ctx context.Context, userID uuid.UUID) ([]sentinelengine_models.UserRole, error) {
	var edges []sentinelengine_models.UserRole
	err := r.conn(ctx).Where("user_id = ?", userID).Find(&edges).Error
	return edges, err
}

// GetUsersAssignedToRole returns the assignment edges for a role within an org.
func (r *sentinelRepo) GetUsersAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID) ([]sentinelengine_models.UserRole, error) {
	var edges []sentinelengine_models.UserRole
	err := r.conn(ctx).
		Where("organization_id = ? AND role_id = ?", organizationID, roleID).
		Find(&edges).Error
	return edges, err
}

// IsRoleAssignedToUser reports whether a user has a role.
func (r *sentinelRepo) IsRoleAssignedToUser(ctx context.Context, userID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).Error
	return count > 0, err
}

// ReplaceRolesAssignedToUser sets a user's roles to exactly roleIDs, in one tx.
func (r *sentinelRepo) ReplaceRolesAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, roleIDs []uuid.UUID) error {
	return r.conn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("organization_id = ? AND user_id = ?", organizationID, userID).
			Delete(&sentinelengine_models.UserRole{}).Error; err != nil {
			return err
		}
		for _, roleID := range roleIDs {
			edge := &sentinelengine_models.UserRole{
				ID:             uuid.New(),
				OrganizationID: organizationID,
				UserID:         userID,
				RoleID:         roleID,
			}
			if err := tx.Create(edge).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
