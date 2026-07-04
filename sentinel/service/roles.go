package sentinelservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateRole implements [RolesService]: create a role and optionally attach an
// initial set of existing policies, in one transaction.
func (s *sentinelService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, errors.New("role name is required")
	}

	exists, err := s.repo.RoleExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("a role named %q already exists", name)
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
		for _, policyID := range in.PolicyIDs {
			policy, err := s.repo.GetPolicyByID(txCtx, policyID)
			if err != nil {
				return fmt.Errorf("policy %s: %w", policyID, err)
			}
			if policy.OrganizationID != orgID {
				return fmt.Errorf("policy %s not found in this organization", policyID)
			}
			edge := &sentinelengine_models.RolePolicies{
				OrganizationID:   orgID,
				RoleID:           role.ID,
				SentinelPolicyID: policyID,
			}
			if err := s.repo.AssignPolicyToRole(txCtx, edge); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return uuid.Nil, err
	}

	s.invalidateEngine(orgID)
	return role.ID, nil
}

// ListRoles implements [RolesService].
func (s *sentinelService) ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	roles, err := s.repo.ListRolesByOrganizationID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToRoleViews(roles), nil
}

// GetRolePrivileges implements [RolesService]: the role plus every policy
// attached to it.
func (s *sentinelService) GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, err
	}
	if role.OrganizationID != orgID {
		return sentinelDtos.RolePrivilegesView{}, errors.New("role not found in this organization")
	}

	policies, err := s.repo.ListPoliciesAssignedToRole(ctx, orgID, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, err
	}

	return sentinelDtos.RolePrivilegesView{
		Role:     sentinelDtos.ToRoleView(&role),
		Policies: sentinelDtos.ToPolicyViews(policies),
	}, nil
}

// GetEmployeeRole implements [SentinelServiceInterface]: the role assigned to
// an employee within an org (first one, if several).
func (s *sentinelService) GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error) {
	edges, err := s.repo.GetRolesAssignedToUser(ctx, emplID)
	if err != nil {
		return sentinelDtos.RoleView{}, err
	}
	for _, edge := range edges {
		if edge.OrganizationID != orgID {
			continue
		}
		role, err := s.repo.GetRoleByID(ctx, edge.RoleID)
		if err != nil {
			return sentinelDtos.RoleView{}, err
		}
		return sentinelDtos.ToRoleView(&role), nil
	}
	return sentinelDtos.RoleView{}, errors.New("employee has no role in this organization")
}
