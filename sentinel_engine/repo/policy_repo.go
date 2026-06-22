package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// CreatePolicy persists a single Cedar policy row. The caller is responsible for
// validating that policy.Policy compiles before calling this. Joins an ambient
// transaction when one is present in the context.
func (r *sentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	return r.conn(ctx).Create(policy).Error
}
