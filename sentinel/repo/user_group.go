package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AssignUserToGroup adds a user to a group. Idempotent: if the user is already a
// member it is a no-op.
func (r *sentinelRepo) AssignUserToGroup(ctx context.Context, organizationID, userID, groupID uuid.UUID) error {
	assigned, err := r.IsUserAssignedToGroup(ctx, userID, groupID)
	if err != nil {
		return err
	}
	if assigned {
		return nil
	}
	edge := &sentinelengine_models.UserGroup{
		ID:             uuid.New(),
		OrganizationID: organizationID,
		UserID:         userID,
		GroupID:        groupID,
	}
	return r.conn(ctx).Create(edge).Error
}

// RemoveUserFromGroup removes a membership (hard delete — no soft-delete column).
func (r *sentinelRepo) RemoveUserFromGroup(ctx context.Context, organizationID, userID, groupID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND user_id = ? AND group_id = ?", organizationID, userID, groupID).
		Delete(&sentinelengine_models.UserGroup{}).Error
}

// GetGroupsAssignedToUser returns the membership edges for a user.
func (r *sentinelRepo) GetGroupsAssignedToUser(ctx context.Context, userID uuid.UUID) ([]sentinelengine_models.UserGroup, error) {
	var edges []sentinelengine_models.UserGroup
	err := r.conn(ctx).Where("user_id = ?", userID).Find(&edges).Error
	return edges, err
}

// GetUsersAssignedToGroup returns the membership edges for a group within an org.
func (r *sentinelRepo) GetUsersAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID) ([]sentinelengine_models.UserGroup, error) {
	var edges []sentinelengine_models.UserGroup
	err := r.conn(ctx).
		Where("organization_id = ? AND group_id = ?", organizationID, groupID).
		Find(&edges).Error
	return edges, err
}

// IsUserAssignedToGroup reports whether a user is a member of a group.
func (r *sentinelRepo) IsUserAssignedToGroup(ctx context.Context, userID, groupID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.UserGroup{}).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Count(&count).Error
	return count > 0, err
}

// ReplaceGroupsAssignedToUser sets a user's groups to exactly groupIDs, in one tx.
func (r *sentinelRepo) ReplaceGroupsAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, groupIDs []uuid.UUID) error {
	return r.conn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("organization_id = ? AND user_id = ?", organizationID, userID).
			Delete(&sentinelengine_models.UserGroup{}).Error; err != nil {
			return err
		}
		for _, groupID := range groupIDs {
			edge := &sentinelengine_models.UserGroup{
				ID:             uuid.New(),
				OrganizationID: organizationID,
				UserID:         userID,
				GroupID:        groupID,
			}
			if err := tx.Create(edge).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
