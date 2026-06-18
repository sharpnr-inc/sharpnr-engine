package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// CreateRole inserts a new role. GORM fills in timestamps on success.
func (r *sentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	return r.conn(ctx).WithContext(ctx).Create(role).Error
}

// AssignRole links an employee to a role.
func (r *sentinelRepo) AssignRole(ctx context.Context, assignment *sentinelengine_models.EmployeeRole) error {
	return r.conn(ctx).WithContext(ctx).Create(assignment).Error
}
