package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

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
