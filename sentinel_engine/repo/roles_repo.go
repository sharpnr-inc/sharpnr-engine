package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

func (r *sentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	panic("")
}

func (r *sentinelRepo) AssignUserRole(ctx context.Context, userID, roleID uuid.UUID) error {
	panic("")
}
