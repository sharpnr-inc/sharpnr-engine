package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// This file backs the managed-policy authoring layer and the generator. Writes
// to managed_policies / role_policies are the source of truth; the generator
// reads the attachments here and (re)writes the generated sentinel_policies
// rows that the engine actually evaluates.

// CreateManagedPolicy persists a reusable managed policy. ID is app-generated.
func (r *sentinelRepo) CreateManagedPolicy(ctx context.Context, mp *sentinelengine_models.ManagedPolicy) error {
	if mp.ID == uuid.Nil {
		mp.ID = uuid.New()
	}
	return r.conn(ctx).Create(mp).Error
}

// AttachRolePolicy inserts the (role, managed policy) attachment edge.
func (r *sentinelRepo) AttachRolePolicy(ctx context.Context, edge *sentinelengine_models.RolePolicy) error {
	return r.conn(ctx).Create(edge).Error
}

// DetachRolePolicy removes an attachment (hard delete; the pair can be re-added).
func (r *sentinelRepo) DetachRolePolicy(ctx context.Context, orgID, roleID, managedPolicyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND role_id = ? AND managed_policy_id = ?", orgID, roleID, managedPolicyID).
		Delete(&sentinelengine_models.RolePolicy{}).Error
}

// ListManagedPolicies returns the managed policies visible to an org: the org's
// own policies plus the platform "system" policies (organization_id IS NULL).
// This is the catalog the UI shows in the "attach policies" picker.
func (r *sentinelRepo) ListManagedPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error) {
	var policies []*sentinelengine_models.ManagedPolicy
	err := r.conn(ctx).
		Where("organization_id = ? OR organization_id IS NULL", orgID).
		Order("system DESC, name ASC").
		Find(&policies).Error
	return policies, err
}

// GetManagedPolicy loads one managed policy visible to the org (its own or a
// system policy). Used to validate an attachment target belongs to the tenant.
func (r *sentinelRepo) GetManagedPolicy(ctx context.Context, orgID, id uuid.UUID) (*sentinelengine_models.ManagedPolicy, error) {
	var mp sentinelengine_models.ManagedPolicy
	err := r.conn(ctx).
		Where("id = ? AND (organization_id = ? OR organization_id IS NULL)", id, orgID).
		First(&mp).Error
	if err != nil {
		return nil, err
	}
	return &mp, nil
}

// ListRoles returns all live roles for an org (the Add-User / role-list source).
func (r *sentinelRepo) ListRoles(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Role, error) {
	var roles []*sentinelengine_models.Role
	err := r.conn(ctx).
		Where("organization_id = ?", orgID).
		Order("name ASC").
		Find(&roles).Error
	return roles, err
}

// ListAttachedManagedPolicies returns the managed policies attached to a role.
func (r *sentinelRepo) ListAttachedManagedPolicies(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error) {
	var policies []*sentinelengine_models.ManagedPolicy
	err := r.conn(ctx).
		Table("sentinel_engine.managed_policies AS mp").
		Joins("JOIN sentinel_engine.role_policies AS rp ON rp.managed_policy_id = mp.id").
		Where("rp.organization_id = ? AND rp.role_id = ? AND mp.deleted_at IS NULL", orgID, roleID).
		Find(&policies).Error
	return policies, err
}

// DeleteGeneratedPoliciesByRole hard-deletes the generated sentinel_policies the
// generator owns for a role (source_role_id = roleID). Hand-authored policies
// (source_role_id NULL) are never matched, so they survive regeneration.
func (r *sentinelRepo) DeleteGeneratedPoliciesByRole(ctx context.Context, orgID, roleID uuid.UUID) error {
	return r.conn(ctx).
		Unscoped().
		Where("organization_id = ? AND source_role_id = ?", orgID, roleID).
		Delete(&sentinelengine_models.SentinelPolicy{}).Error
}
