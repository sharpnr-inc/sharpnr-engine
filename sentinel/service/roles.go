package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelServiceErrs "sharpnr-engine/sentinel/errs"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateRole implements [RolesService]: create a role and optionally register an
// initial set of existing policies against it, in one transaction.
func (s *sentinelService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, sentinelServiceErrs.RoleNameRequiredErr
	}

	exists, err := s.repo.RoleExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("%w: %q", sentinelServiceErrs.RoleAlreadyExistsErr, name)
	}

	role := &sentinelengine_models.Role{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Description:    in.Description,
	}

	if err := s.withTx(ctx, func(txCtx context.Context) error {

		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return err
		}
		// for _, policyID := range in.PolicyIDs {
		// 	if _, err := s.repo.GetPolicy(txCtx, orgID, policyID); err != nil {
		// 		return fmt.Errorf("%w: %s", sentinelServiceErrs.PolicyNotFoundInOrgErr, policyID)
		// 	}
		// 	if err := s.repo.AttachPolicyToRole(txCtx, orgID, role.ID, policyID); err != nil {
		// 		return err
		// 	}
		// }
		return nil
	}); err != nil {
		return uuid.Nil, err
	}

	s.invalidateEngine(orgID)
	return role.ID, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListRoles implements [RolesService].
func (s *sentinelService) ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	roles, err := s.repo.ListRoles(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToRoleViews(roles), nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetRolePrivileges implements [RolesService]: the role plus every policy
// registered against it. Includes disabled policies — an admin inspecting a role
// wants to see them; the engine is what filters on enabled.
func (s *sentinelService) GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	role, err := s.repo.GetRole(ctx, orgID, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, fmt.Errorf("%w: %s", sentinelServiceErrs.RoleNotFoundInOrgErr, roleID)
	}

	policies, err := s.repo.ListPoliciesOfRole(ctx, orgID, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, err
	}

	return sentinelDtos.RolePrivilegesView{
		Role:     sentinelDtos.ToRoleView(&role),
		Policies: sentinelDtos.ToPolicyViews(policies),
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetEmployeeRole implements [SentinelServiceInterface]: the role an employee
// holds in an org.
//
// One org-scoped query returning entities. It used to fetch the employee's role
// edges across EVERY org and filter in Go, then re-fetch each role by id.
//
// The schema allows several roles per employee; this returns the most recently
// created one. Callers wanting all of them should use the repo's GetRolesOfUser.
func (s *sentinelService) GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error) {
	roles, err := s.repo.GetRolesOfUser(ctx, orgID, emplID)
	if err != nil {
		return sentinelDtos.RoleView{}, err
	}
	if len(roles) == 0 {
		// The caller decides whether an employee without a role is an error or
		// simply an unassigned user.
		return sentinelDtos.RoleView{}, sentinelServiceErrs.EmployeeRoleNotFoundErr
	}
	return sentinelDtos.ToRoleView(roles[0]), nil
}
