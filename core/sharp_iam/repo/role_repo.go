package sharpIam_repo

import (
	"context"
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

	"github.com/google/uuid"
)

func (r *sharpIamRepository) CreateRole(ctx context.Context, role *sharp_iam_models.SharpIamRoles) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *sharpIamRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (*sharp_iam_models.SharpIamRoles, error) {
	var role sharp_iam_models.SharpIamRoles
	if err := r.db.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *sharpIamRepository) UpdateRole(ctx context.Context, role *sharp_iam_models.SharpIamRoles) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *sharpIamRepository) DeleteRole(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&sharp_iam_models.SharpIamRoles{}, "id = ?", id).Error
}

func (r *sharpIamRepository) ListRoles(ctx context.Context, orgID uuid.UUID, domain string) ([]*sharp_iam_models.SharpIamRoles, error) {
	var roles []*sharp_iam_models.SharpIamRoles
	if err := r.db.WithContext(ctx).Where("org_id = ? AND domain = ?", orgID, domain).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
