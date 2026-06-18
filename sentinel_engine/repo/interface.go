package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// SentinelRepository is the data-access layer for the authorization tables.
// Each method participates in an ambient transaction when one is present in the
// context (see dbtx), so a service can run several writes atomically.
type SentinelRepository interface {
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error
	CreatePermission(ctx context.Context, permission *sentinelengine_models.Permission) error
	AssignRole(ctx context.Context, assignment *sentinelengine_models.EmployeeRole) error
}
