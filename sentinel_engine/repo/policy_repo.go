package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// CreatePermission inserts a single permission row for a role.
func (r *sentinelRepo) CreatePermission(ctx context.Context, permission *sentinelengine_models.Permission) error {
	return r.conn(ctx).WithContext(ctx).Create(permission).Error
}
