package sentinel_repo

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateRole persists a new role. The ID is application-generated; if the caller
// left it unset we assign one here so role.ID is valid immediately.
func (r *sentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	if role.ID == uuid.Nil {
		role.ID = uuid.New()
	}
	return r.conn(ctx).Create(role).Error
}

// GetRoleByID loads a role by its primary key.
func (r *sentinelRepo) GetRoleByID(ctx context.Context, roleID uuid.UUID) (sentinelengine_models.Role, error) {
	var role sentinelengine_models.Role
	err := r.conn(ctx).Where("id = ?", roleID).First(&role).Error
	return role, err
}

// GetRoleByName loads a role by its (org, name) — name is the role's unique
// per-org identifier (and Cedar entity id).
func (r *sentinelRepo) GetRoleByName(ctx context.Context, organizationID uuid.UUID, roleName string) (sentinelengine_models.Role, error) {
	var role sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, roleName).
		First(&role).Error
	return role, err
}

// ListRolesByOrganizationID returns every role in an org, newest first.
func (r *sentinelRepo) ListRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Role, error) {
	var roles []*sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("created_at DESC").
		Find(&roles).Error
	return roles, err
}

// UpdateRoleDetails applies a partial update (name and/or description). Nil
// fields are left unchanged; a no-op input is a no-op.
func (r *sentinelRepo) UpdateRoleDetails(ctx context.Context, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error {
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
		Where("id = ?", roleID).
		Updates(fields).Error
}

// DeleteRole soft-deletes a role (the model carries gorm.DeletedAt).
func (r *sentinelRepo) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	return r.conn(ctx).Where("id = ?", roleID).Delete(&sentinelengine_models.Role{}).Error
}

// RoleExists reports whether a role with that name already exists in the org.
func (r *sentinelRepo) RoleExists(ctx context.Context, organizationID uuid.UUID, roleName string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.Role{}).
		Where("organization_id = ? AND name = ?", organizationID, roleName).
		Count(&count).Error
	return count > 0, err
}
