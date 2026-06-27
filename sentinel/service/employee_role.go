package sentinelservice

import (
	"context"
	"fmt"

	sentinelDtos "sharpnr-engine/sentinel/dtos"

	"github.com/google/uuid"
)

// GetEmployeeRole resolves the role assigned to an employee within an org:
// the user_roles edge -> the role's details. Returns the role view the UI shows
// for "this employee's role".
func (s *sentinelService) GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error) {
	ur, err := s.repo.GetUserRole(ctx, orgID, emplID)
	if err != nil {
		return sentinelDtos.RoleView{}, fmt.Errorf("get role assignment for employee %s: %w", emplID, err)
	}

	role, err := s.repo.GetRoleByID(ctx, ur.RoleID)
	if err != nil {
		return sentinelDtos.RoleView{}, fmt.Errorf("get role %s: %w", ur.RoleID, err)
	}

	return sentinelDtos.ToRoleView(&role), nil
}
