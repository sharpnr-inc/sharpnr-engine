package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// Groups are the second distribution mechanism, alongside roles. Like roles
// they are NOT Cedar entities — a policy reaches a user because it is
// registered against a group the user belongs to.
//
// Groups nest through ParentGroupID. Note that nesting is presentation only
// today: GetAllPoliciesOfUser does not walk the tree, so a policy registered
// against a parent group does not reach members of its children.

// CreateGroup persists a group. ID is application-generated if unset.
func (r *sentinelRepo) CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error {
	if group.ID == uuid.Nil {
		group.ID = uuid.New()
	}
	return r.conn(ctx).Create(group).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetGroup loads one of the org's groups. A group belonging to another org is
// reported as not found, never returned.
func (r *sentinelRepo) GetGroup(ctx context.Context, organizationID, groupID uuid.UUID) (sentinelengine_models.Group, error) {
	var group sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, groupID).
		First(&group).Error
	return group, err
}

/* ---------------------------------- <***> --------------------------------- */

// GetGroupByName loads a group by its (org, name) pair.
//
// NOTE: unlike roles, groups have no unique index on (org, name) — 000039
// dropped the slug without adding one. Duplicate names are therefore possible
// and this returns the first match; GroupExists is what the service checks
// before creating.
func (r *sentinelRepo) GetGroupByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Group, error) {
	var group sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, name).
		First(&group).Error
	return group, err
}

/* ---------------------------------- <***> --------------------------------- */

// ListGroups returns every group in an org, newest first.
func (r *sentinelRepo) ListGroups(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	var groups []*sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("created_at DESC").
		Find(&groups).Error
	return groups, err
}

/* ---------------------------------- <***> --------------------------------- */

// ListChildGroups returns the direct children of a group — one level, not the
// whole subtree.
func (r *sentinelRepo) ListChildGroups(ctx context.Context, organizationID, parentGroupID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	var groups []*sentinelengine_models.Group
	err := r.conn(ctx).
		Where("organization_id = ? AND parent_group_id = ?", organizationID, parentGroupID).
		Order("created_at DESC").
		Find(&groups).Error
	return groups, err
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateGroup applies a partial update (name, description and/or parent). Nil
// fields are left unchanged; a no-op input is a no-op.
//
// ParentGroupID is a *uuid.UUID on the model, so "clear the parent" and "leave
// it alone" are the same value here. Re-parenting to top level needs its own
// method rather than being smuggled through this one.
func (r *sentinelRepo) UpdateGroup(ctx context.Context, organizationID uuid.UUID, group *sentinelengine_models.Group) error {
	if group == nil {
		return nil
	}
	fields := map[string]any{
		"name":        group.Name,
		"description": group.Description,
	}
	if group.ParentGroupID != nil {
		fields["parent_group_id"] = *group.ParentGroupID
	}
	return r.conn(ctx).
		Model(&sentinelengine_models.Group{}).
		Where("organization_id = ? AND id = ?", organizationID, group.ID).
		Updates(fields).Error
}

/* ---------------------------------- <***> --------------------------------- */

// DeleteGroup soft-deletes a group.
//
// NOTE: groups_sentinel_policies cascades from groups(id), but only on a hard
// delete — a soft delete leaves the policy registrations in place. Child groups
// are not touched either and will point at a deleted parent, so the service
// must decide what happens to the subtree.
func (r *sentinelRepo) DeleteGroup(ctx context.Context, organizationID, groupID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, groupID).
		Delete(&sentinelengine_models.Group{}).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GroupExists reports whether a group with that name already exists in the org.
func (r *sentinelRepo) GroupExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.Group{}).
		Where("organization_id = ? AND name = ?", organizationID, name).
		Count(&count).Error
	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */
