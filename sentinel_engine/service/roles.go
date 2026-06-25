package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateRole creates a role and, optionally, attaches an initial set of existing
// policies to it — all in one transaction. The slug is the role's stable Cedar
// entity id; policies reference it as Role::"<slug>", so it is derived from the
// name when the caller doesn't supply one.
func (s *sentinelService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, fmt.Errorf("role name is required")
	}

	slug := strings.TrimSpace(in.Slug)
	if slug == "" {
		slug = slugify(name)
	}

	role := &sentinelengine_models.Role{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Slug:           slug,
		Name:           name,
		Description:    in.Description,
	}

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return fmt.Errorf("create role: %w", err)
		}
		for _, policyID := range in.PolicyIDs {
			rp := &sentinelengine_models.RolePolicies{
				ID:               uuid.New(),
				OrganizationID:   orgID,
				RoleID:           role.ID,
				SentinelPolicyID: policyID,
			}
			if err := s.repo.CreateRolePolicy(txCtx, rp); err != nil {
				return fmt.Errorf("attach policy %s to role: %w", policyID, err)
			}
		}
		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// ListRoles returns every role defined for an org.
func (s *sentinelService) ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	roles, err := s.repo.ListRolesByOrgID(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("list roles: %w", err)
	}
	return sentinelDtos.ToRoleViews(roles), nil
}

// GetRolePrivileges returns a role together with the policies attached to it.
func (s *sentinelService) GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, fmt.Errorf("get role %s: %w", roleID, err)
	}
	if role.OrganizationID != orgID {
		return sentinelDtos.RolePrivilegesView{}, fmt.Errorf("role %s does not belong to org %s", roleID, orgID)
	}

	policies, err := s.repo.ListPoliciesByRoleID(ctx, orgID, roleID)
	if err != nil {
		return sentinelDtos.RolePrivilegesView{}, fmt.Errorf("list policies for role %s: %w", roleID, err)
	}

	return sentinelDtos.RolePrivilegesView{
		Role:     sentinelDtos.ToRoleView(&role),
		Policies: sentinelDtos.ToPolicyViews(policies),
	}, nil
}

// slugify turns a human role/group name into a stable, lowercase, dash-separated
// slug usable as a Cedar entity id.
func slugify(name string) string {
	var b strings.Builder
	prevDash := false
	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			b.WriteRune(r)
			prevDash = false
		default:
			if !prevDash && b.Len() > 0 {
				b.WriteByte('-')
				prevDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
