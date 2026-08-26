package sentinel_repo

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// Roles are a distribution mechanism, not a Cedar entity: a policy reaches a
// user because it is registered against a role the user holds, never because
// its text names the role. See attachments.go (role -> policy) and
// memberships.go (user -> role).
//
// Name is unique per org. It is a human label only — nothing embeds it in
// policy text — but the uniqueness index makes it a stable handle for the UI.

// CreateRole persists a role. ID is application-generated if unset.
func (r *sentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	if role.ID == uuid.Nil {
		role.ID = uuid.New()
	}
	return r.conn(ctx).Create(role).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetRole loads one of the org's roles. A role belonging to another org is
// reported as not found, never returned.
func (r *sentinelRepo) GetRole(ctx context.Context, organizationID, roleID uuid.UUID) (sentinelengine_models.Role, error) {
	var role sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, roleID).
		First(&role).Error
	return role, err
}

/* ---------------------------------- <***> --------------------------------- */

// GetRoleByName loads a role by its (org, name) pair.
func (r *sentinelRepo) GetRoleByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Role, error) {
	var role sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, name).
		First(&role).Error
	return role, err
}

/* ---------------------------------- <***> --------------------------------- */

// ListRoles returns every role in an org, newest first.
func (r *sentinelRepo) ListRoles(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Role, error) {
	var roles []*sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("created_at DESC").
		Find(&roles).Error
	return roles, err
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateRole applies a partial update (name and/or description). Nil fields are
// left unchanged; a no-op input is a no-op.
func (r *sentinelRepo) UpdateRole(ctx context.Context, organizationID, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error {
	if updates == nil {
		return nil
	}
	fields := map[string]any{}
	if updates.Name != nil {
		fields["name"] = *updates.Name
	}
	if updates.Description != nil {
		fields["description"] = *updates.Description
	}
	if len(fields) == 0 {
		return nil
	}
	return r.conn(ctx).
		Model(&sentinelengine_models.Role{}).
		Where("organization_id = ? AND id = ?", organizationID, roleID).
		Updates(fields).Error
}

/* ---------------------------------- <***> --------------------------------- */

// DeleteRole soft-deletes a role.
//
// NOTE: roles.id has no FK from role_policies or user_roles (the repo-wide
// convention), so nothing cascades. The service must detach the role's policies
// and remove its memberships itself, or they are left dangling.
func (r *sentinelRepo) DeleteRole(ctx context.Context, organizationID, roleID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, roleID).
		Delete(&sentinelengine_models.Role{}).Error
}

/* ---------------------------------- <***> --------------------------------- */

// RoleExists reports whether a role with that name already exists in the org.
func (r *sentinelRepo) RoleExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.Role{}).
		Where("organization_id = ? AND name = ?", organizationID, name).
		Count(&count).Error
	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */
