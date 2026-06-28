package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateGroup persists a new group. ID is application-generated if unset.
func (r *sentinelRepo) CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error {
	if group.ID == uuid.Nil {
		group.ID = uuid.New()
	}
	return r.conn(ctx).Create(group).Error
}

// GetGroupByID loads a group by its primary key.
func (r *sentinelRepo) GetGroupByID(ctx context.Context, groupID uuid.UUID) (sentinelengine_models.Group, error) {
	var group sentinelengine_models.Group
	err := r.conn(ctx).Where("id = ?", groupID).First(&group).Error
	return group, err
}

// GetGroupByName loads a group by its (org, name) — name identifies a group
// within an org (slug was removed in migration 000039).
func (r *sentinelRepo) GetGroupByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Group, error) {
	var group sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, name).
		First(&group).Error
	return group, err
}

// ListGroupsByOrganizationID returns every group in an org, oldest first so
// parents tend to precede children.
func (r *sentinelRepo) ListGroupsByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	var groups []*sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("created_at ASC").
		Find(&groups).Error
	return groups, err
}

// ListChildGroups returns the direct children of a group.
func (r *sentinelRepo) ListChildGroups(ctx context.Context, parentGroupID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	var groups []*sentinelengine_models.Group
	err := r.conn(ctx).
		Where("parent_group_id = ?", parentGroupID).
		Order("created_at ASC").
		Find(&groups).Error
	return groups, err
}

// UpdateGroup persists all columns of the given group by its primary key.
func (r *sentinelRepo) UpdateGroup(ctx context.Context, group *sentinelengine_models.Group) error {
	return r.conn(ctx).Save(group).Error
}

// DeleteGroup soft-deletes a group.
func (r *sentinelRepo) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	return r.conn(ctx).Where("id = ?", groupID).Delete(&sentinelengine_models.Group{}).Error
}

// GroupExists reports whether a group with that name already exists in the org.
func (r *sentinelRepo) GroupExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.Group{}).
		Where("organization_id = ? AND name = ?", organizationID, name).
		Count(&count).Error
	return count > 0, err
}
