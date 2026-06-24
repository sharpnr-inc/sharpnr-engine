package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// CreateGroup persists a group. The ID is application-generated; if the caller
// left it unset we assign one here so group.ID is valid immediately (e.g. to add
// members in the same transaction).
func (r *sentinelRepo) CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error {
	if group.ID == uuid.Nil {
		group.ID = uuid.New()
	}
	return r.conn(ctx).Create(group).Error
}
