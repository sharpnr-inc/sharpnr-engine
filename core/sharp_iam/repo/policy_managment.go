package sharpIam_repo

import (
	"context"
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

	"github.com/google/uuid"
)

func (r *sharpIamRepository) CreatePolicy(ctx context.Context, policy *sharp_iam_models.SharpIamPolicies) error {
	return r.db.WithContext(ctx).Create(policy).Error
}

func (r *sharpIamRepository) GetPolicyByID(ctx context.Context, id uuid.UUID) (*sharp_iam_models.SharpIamPolicies, error) {
	panic("")
}

func (r *sharpIamRepository) UpdatePolicy(ctx context.Context, policy *sharp_iam_models.SharpIamPolicies) error {
	panic("")
}

func (r *sharpIamRepository) DeletePolicy(ctx context.Context, id uuid.UUID) error {
	panic("")
}

func (r *sharpIamRepository) ListPolicies(ctx context.Context, orgID uuid.UUID, domain string) ([]*sharp_iam_models.SharpIamPolicies, error) {
	panic("")
}
