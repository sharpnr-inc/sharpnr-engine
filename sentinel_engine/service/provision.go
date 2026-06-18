package sentinelservice

import (
	"context"
	"fmt"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProvisionSuperAdmin sets up authorization for a brand-new organization. It
// creates, in one transaction:
//
//  1. a system role "Super Admin" for the org,
//  2. one wildcard permission on that role (every action, every resource type,
//     every property in the org), and
//  3. an assignment linking the founding employee to that role.
//
// After this runs, that employee passes every authorization check inside their
// org. It is called once, during signup, right after the default employee is
// created.
func (s *sentinelService) ProvisionSuperAdmin(ctx context.Context, orgID, employeeID uuid.UUID) (uuid.UUID, error) {
	roleID := uuid.New()

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. the role
		role := &sentinelengine_models.Role{
			ID:             roleID,
			OrganizationID: orgID,
			Name:           "Super Admin",
			Description:    "Full access to everything in the organization.",
			IsSystem:       true,
		}
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return fmt.Errorf("create super admin role: %w", err)
		}

		// 2. the wildcard permission (PropertyID nil = all properties,
		//    ResourceType "*" = all modules, all CRUD = full control)
		permission := &sentinelengine_models.Permission{
			ID:           uuid.New(),
			RoleID:       roleID,
			PropertyID:   nil,
			ResourceType: "*",
			CanView:      true,
			CanCreate:    true,
			CanEdit:      true,
			CanDelete:    true,
		}
		if err := s.repo.CreatePermission(txCtx, permission); err != nil {
			return fmt.Errorf("create super admin permission: %w", err)
		}

		// 3. assign the role to the founding employee
		assignment := &sentinelengine_models.EmployeeRole{
			ID:             uuid.New(),
			EmployeeID:     employeeID,
			RoleID:         roleID,
			OrganizationID: orgID,
		}
		if err := s.repo.AssignRole(txCtx, assignment); err != nil {
			return fmt.Errorf("assign super admin role to employee: %w", err)
		}

		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return roleID, nil
}
