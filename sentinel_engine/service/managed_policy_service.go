package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file is the managed-policy authoring + display surface the orchestrator
// exposes. Mutations follow the same shape: write the structured source of
// truth in one committed transaction, THEN regenerate the role's Cedar. The
// regenerate step runs in its own transaction and invalidates the org cache
// only after commit, so it must run on a context with no ambient tx — hence we
// never wrap "write + regenerate" in a single outer transaction.

// Catalog returns the full permission catalog (the picker's data source).
func (s *sentinelService) Catalog() []sentinel_actions.Permission {
	return sentinel_actions.Catalog()
}

// ListManagedPolicies returns the managed policies an org can attach (its own +
// system), each with its actions expanded into display permissions.
func (s *sentinelService) ListManagedPolicies(ctx context.Context, orgID uuid.UUID) ([]ManagedPolicyView, error) {
	rows, err := s.repo.ListManagedPolicies(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("list managed policies: %w", err)
	}
	views := make([]ManagedPolicyView, 0, len(rows))
	for _, mp := range rows {
		views = append(views, toManagedPolicyView(mp))
	}
	return views, nil
}

// CreateManagedPolicy authors a new org-scoped managed policy. Actions are
// validated against the catalog so we never store (and later generate) a
// reference to an action id that doesn't exist.
func (s *sentinelService) CreateManagedPolicy(ctx context.Context, orgID uuid.UUID, in CreateManagedPolicyInput) (uuid.UUID, error) {
	if err := validateActions(in.Actions); err != nil {
		return uuid.Nil, err
	}

	effect := sentinelengine_models.EffectPermit
	if strings.EqualFold(in.Effect, string(sentinelengine_models.EffectForbid)) {
		effect = sentinelengine_models.EffectForbid
	}

	mp := &sentinelengine_models.ManagedPolicy{
		ID:             uuid.New(),
		OrganizationID: &orgID, // org-authored (never system via this path)
		Name:           in.Name,
		Description:    in.Description,
		Actions:        in.Actions,
		Effect:         effect,
		ResourceType:   in.ResourceType,
		System:         false,
	}
	if err := s.repo.CreateManagedPolicy(ctx, mp); err != nil {
		return uuid.Nil, fmt.Errorf("create managed policy: %w", err)
	}
	return mp.ID, nil
}

// ListRoles returns the org's roles (e.g. for the Add-User dropdown).
func (s *sentinelService) ListRoles(ctx context.Context, orgID uuid.UUID) ([]RoleView, error) {
	rows, err := s.repo.ListRoles(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("list roles: %w", err)
	}
	views := make([]RoleView, 0, len(rows))
	for _, r := range rows {
		views = append(views, toRoleView(r))
	}
	return views, nil
}

// CreateRole creates a role and optionally attaches initial managed policies,
// then generates its Cedar. The create + attach run in one committed
// transaction; regeneration runs after, so its cache invalidation is durable.
func (s *sentinelService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in CreateRoleInput) (uuid.UUID, error) {
	roleID := uuid.New()

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		role := &sentinelengine_models.Role{
			ID:             roleID,
			OrganizationID: orgID,
			Slug:           slugify(in.Name),
			Name:           in.Name,
			Description:    in.Description,
		}
		if role.Slug == "" {
			return fmt.Errorf("role name %q produces an empty slug", in.Name)
		}
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return fmt.Errorf("create role: %w", err)
		}

		for _, pid := range in.ManagedPolicyIDs {
			if err := s.attachOne(txCtx, orgID, roleID, pid, emplID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if len(in.ManagedPolicyIDs) > 0 {
		if err := s.RegenerateRolePolicies(ctx, orgID, roleID); err != nil {
			return uuid.Nil, fmt.Errorf("generate policies for new role: %w", err)
		}
	}
	return roleID, nil
}

// AttachPolicies attaches managed policies to a role and regenerates its Cedar.
func (s *sentinelService) AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		for _, pid := range policyIDs {
			if err := s.attachOne(txCtx, orgID, roleID, pid, emplID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return s.RegenerateRolePolicies(ctx, orgID, roleID)
}

// DetachPolicy removes one attachment and regenerates the role's Cedar.
func (s *sentinelService) DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	if err := s.repo.DetachRolePolicy(ctx, orgID, roleID, policyID); err != nil {
		return fmt.Errorf("detach policy: %w", err)
	}
	return s.RegenerateRolePolicies(ctx, orgID, roleID)
}

// GetRolePrivileges returns the role plus its attached policies and the deduped
// union of permissions they grant — the role detail view.
func (s *sentinelService) GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (RolePrivilegesView, error) {
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return RolePrivilegesView{}, fmt.Errorf("load role %s: %w", roleID, err)
	}
	if role.OrganizationID != orgID {
		return RolePrivilegesView{}, fmt.Errorf("role %s does not belong to org %s", roleID, orgID)
	}

	attached, err := s.repo.ListAttachedManagedPolicies(ctx, orgID, roleID)
	if err != nil {
		return RolePrivilegesView{}, fmt.Errorf("load attached policies: %w", err)
	}

	policies := make([]ManagedPolicyView, 0, len(attached))
	seen := map[sentinel_actions.Name]struct{}{}
	effective := make([]sentinel_actions.Permission, 0)
	allActions := false
	for _, mp := range attached {
		view := toManagedPolicyView(mp)
		policies = append(policies, view)
		// Effective = union of permit grants; forbid policies subtract at
		// evaluation time but are not unioned into the "can do" display set.
		if mp.Effect == sentinelengine_models.EffectForbid {
			continue
		}
		if mp.AllActions {
			allActions = true // a permit-all attachment grants the whole catalog
			continue
		}
		for _, p := range view.Permissions {
			if _, dup := seen[p.Key]; dup {
				continue
			}
			seen[p.Key] = struct{}{}
			effective = append(effective, p)
		}
	}

	// An attached "Full Access" policy means the role effectively has every
	// permission in the catalog, so surface the whole catalog rather than the
	// (empty) action list of the wildcard policy.
	if allActions {
		effective = sentinel_actions.Catalog()
	}

	return RolePrivilegesView{
		Role:                 toRoleView(&role),
		Policies:             policies,
		EffectivePermissions: effective,
	}, nil
}

// attachOne validates a managed policy is visible to the org, then inserts the
// (role, policy) edge. Must run inside a transaction (txCtx).
func (s *sentinelService) attachOne(txCtx context.Context, orgID, roleID, policyID, emplID uuid.UUID) error {
	if _, err := s.repo.GetManagedPolicy(txCtx, orgID, policyID); err != nil {
		return fmt.Errorf("managed policy %s not available to org: %w", policyID, err)
	}
	edge := &sentinelengine_models.RolePolicy{
		OrganizationID:  orgID,
		RoleID:          roleID,
		ManagedPolicyID: policyID,
		AttachedBy:      emplID,
	}
	if err := s.repo.AttachRolePolicy(txCtx, edge); err != nil {
		return fmt.Errorf("attach policy %s to role %s: %w", policyID, roleID, err)
	}
	return nil
}

// validateActions rejects any action key not present in the catalog.
func validateActions(actions []string) error {
	for _, a := range actions {
		if !sentinel_actions.IsValid(sentinel_actions.Name(a)) {
			return fmt.Errorf("unknown action %q (not in permission catalog)", a)
		}
	}
	return nil
}

// slugify turns a role name into a stable kebab-case Cedar entity id.
func slugify(name string) string {
	var b strings.Builder
	lastDash := false
	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			b.WriteRune(r)
			lastDash = false
		default:
			if b.Len() > 0 && !lastDash {
				b.WriteByte('-')
				lastDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
