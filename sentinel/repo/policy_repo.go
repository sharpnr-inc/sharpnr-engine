package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreatePolicy persists a single Cedar policy row. The caller is responsible for
// validating that policy.Policy compiles before calling this. Joins an ambient
// transaction when one is present in the context. The ID is application-generated;
// if unset we assign one here.
func (r *sentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	if policy.ID == uuid.Nil {
		policy.ID = uuid.New()
	}
	return r.conn(ctx).Create(policy).Error
}

// GetPolicyByID loads a single policy row by its primary key.
func (r *sentinelRepo) GetPolicyByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.SentinelPolicy, error) {
	var policy sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).Where("id = ?", id).First(&policy).Error
	return policy, err
}

// ListPoliciesByOrgID returns an org's policy rows (enabled or not), newest
// first. An empty scope returns every bucket; a non-empty scope filters to that
// bucket (organization|role|group|user). The authoring UI lists them; the engine
// reads only the enabled subset via ListEnabledPolicies.
func (r *sentinelRepo) ListPoliciesByOrgID(ctx context.Context, orgID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]*sentinelengine_models.SentinelPolicy, error) {
	q := r.conn(ctx).Where("organization_id = ?", orgID)
	if scope != "" {
		q = q.Where("scope = ?", scope)
	}

	var rows []*sentinelengine_models.SentinelPolicy
	err := q.Order("created_at DESC").Find(&rows).Error
	return rows, err
}
