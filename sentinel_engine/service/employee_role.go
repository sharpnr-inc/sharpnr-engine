package sentinelservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// GetEmployeeRole resolves the employee's assigned role within an org.
//
// NOTE (Layer 4): the signature currently returns only error, which is a
// placeholder. The UI needs the role's name/description plus the privileges it
// grants (the catalog labels expanded from the attached managed policies). When
// that DTO lands, change this to return it. The data-access path below is the
// real one it will build on: user_roles edge -> role details.
func (s *sentinelService) GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) error {
	ur, err := s.repo.GetUserRole(ctx, orgID, emplID)
	if err != nil {
		return fmt.Errorf("get role assignment for employee %s: %w", emplID, err)
	}

	role, err := s.repo.GetRoleByID(ctx, ur.RoleID)
	if err != nil {
		return fmt.Errorf("get role %s: %w", ur.RoleID, err)
	}

	// Proven fetch; discard until the DTO return type is added.
	_ = role
	return nil
}
