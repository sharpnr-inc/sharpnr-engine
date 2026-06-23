package sentinelservice

import (
	"context"
	"fmt"
	"sort"
	"strings"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegenerateRolePolicies compiles a role's attached managed policies into Cedar
// and replaces the generated sentinel_policies rows the engine evaluates.
//
// Flow (one transaction):
//  1. resolve the role's slug,
//  2. load its attached managed policies (the authoring source of truth),
//  3. delete the role's previously generated policy rows (by source_role_id),
//  4. compile each attachment to a concrete Cedar policy and insert it, tagged
//     with its source so the next regeneration is idempotent.
//
// After the transaction commits, the org's compiled PolicySet cache is
// invalidated so the next authorization recompiles from the new rows. Call this
// whenever a role's attachments change or an attached managed policy is edited.
func (s *sentinelService) RegenerateRolePolicies(ctx context.Context, orgID, roleID uuid.UUID) error {
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		role, err := s.repo.GetRoleByID(txCtx, roleID)
		if err != nil {
			return fmt.Errorf("load role %s: %w", roleID, err)
		}
		if role.OrganizationID != orgID {
			// Tenant guard: never generate policies for a role in another org.
			return fmt.Errorf("role %s does not belong to org %s", roleID, orgID)
		}

		attached, err := s.repo.ListAttachedManagedPolicies(txCtx, orgID, roleID)
		if err != nil {
			return fmt.Errorf("list attached policies for role %s: %w", roleID, err)
		}

		// Replace, don't merge: drop what we generated last time first.
		if err := s.repo.DeleteGeneratedPoliciesByRole(txCtx, orgID, roleID); err != nil {
			return fmt.Errorf("clear generated policies for role %s: %w", roleID, err)
		}

		for _, mp := range attached {
			text, ok := buildCedarPolicy(role.Slug, mp)
			if !ok {
				continue // managed policy with no actions => nothing to grant
			}

			// Validate before persisting: a row that fails to parse would later
			// fail the whole org closed (the engine fails closed on a bad row).
			var p cedar.Policy
			if err := p.UnmarshalCedar([]byte(text)); err != nil {
				return fmt.Errorf("generated policy for role %q / managed policy %q did not compile: %w", role.Slug, mp.Name, err)
			}

			roleID := roleID
			mpID := mp.ID
			row := &sentinelengine_models.SentinelPolicy{
				ID:                    uuid.New(),
				OrganizationID:        orgID,
				Name:                  fmt.Sprintf("%s · %s", role.Slug, mp.Name),
				Description:           fmt.Sprintf("Generated from managed policy %q attached to role %q.", mp.Name, role.Slug),
				Policy:                text,
				Enabled:               true,
				SourceRoleID:          &roleID,
				SourceManagedPolicyID: &mpID,
			}
			if err := s.repo.CreatePolicy(txCtx, row); err != nil {
				return fmt.Errorf("persist generated policy for role %q: %w", role.Slug, err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Outside the tx: the new rows are committed, so the next read recompiles.
	s.engine.InvalidateOrg(orgID)
	return nil
}

// buildCedarPolicy compiles one managed policy, for one role, into Cedar text.
// It returns ok=false when there is nothing to emit (a non-wildcard policy with
// no actions).
//
// Shapes:
//
//	// explicit action list
//	permit (
//	  principal in Role::"<slug>",
//	  action in [Action::"a", Action::"b"],
//	  resource
//	) when { resource.orgId == principal.orgId };
//
//	// AllActions wildcard (every action, incl. future ones)
//	permit (
//	  principal in Role::"<slug>",
//	  action,
//	  resource
//	) when { resource.orgId == principal.orgId };
//
// The tenant guard keeps the grant scoped to the principal's own org. Actions
// are sorted so regeneration is byte-stable. resource is left unconstrained; the
// action ids are already resource-specific (e.g. "createProperty").
func buildCedarPolicy(roleSlug string, mp *sentinelengine_models.ManagedPolicy) (string, bool) {
	if !mp.AllActions && len(mp.Actions) == 0 {
		return "", false
	}

	effect := "permit"
	if mp.Effect == sentinelengine_models.EffectForbid {
		effect = "forbid"
	}

	var b strings.Builder
	b.WriteString(effect)
	b.WriteString(" (\n  principal in Role::\"")
	b.WriteString(roleSlug)
	b.WriteString("\",\n  ")

	if mp.AllActions {
		// Wildcard: match every action.
		b.WriteString("action")
	} else {
		actions := append([]string(nil), mp.Actions...)
		sort.Strings(actions)
		b.WriteString("action in [")
		for i, a := range actions {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString("Action::\"")
			b.WriteString(a)
			b.WriteString("\"")
		}
		b.WriteString("]")
	}

	b.WriteString(",\n  resource\n) when { resource.orgId == principal.orgId };")

	return b.String(), true
}
