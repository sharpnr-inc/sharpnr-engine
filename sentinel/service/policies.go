package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelServiceErrs "sharpnr-engine/sentinel/errs"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	cedar "github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// validateCedarText checks that the frontend-authored policy text compiles.
// The FE validates with cedar-wasm before sending, but the engine fails closed
// on uncompilable rows, so nothing broken may ever reach the table.
func validateCedarText(text string) error {
	list, err := cedar.NewPolicyListFromBytes("policy.cedar", []byte(text))
	if err != nil {
		return fmt.Errorf("%w: %w", sentinelServiceErrs.PolicyDoesNotCompileErr, err)
	}
	if len(list) == 0 {
		return sentinelServiceErrs.EmptyPolicyStatementsErr
	}
	return nil
}

// CreatePolicy implements [SentinelPolicyService]: persist a raw Cedar policy
// after revalidating it, then invalidate the org's engine cache.
//
// A new policy is attached to NOTHING, so it reaches nobody until someone
// registers it against a principal. That is the safe default the old `scope`
// column could not express — it defaulted to "organization", which meant a
// create call that omitted the field silently granted org-wide.
func (s *sentinelService) CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, sentinelServiceErrs.PolicyNameRequiredErr
	}
	if strings.TrimSpace(in.Policy) == "" {
		return uuid.Nil, sentinelServiceErrs.PolicyTextRequiredErr
	}
	if err := validateCedarText(in.Policy); err != nil {
		return uuid.Nil, err
	}

	// sentinel_policies has no unique index on (org, name), so the check is the
	// only thing keeping names distinct.
	exists, err := s.repo.PolicyExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("%w: %q", sentinelServiceErrs.PolicyAlreadyExistsErr, name)
	}

	policy := &sentinelengine_models.SentinelPolicy{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Description:    in.Description,
		Policy:         in.Policy,
		PolicyJson:     in.PolicyJson,
		Enabled:        in.IsEnabled,
	}
	if err := s.repo.CreatePolicy(ctx, policy); err != nil {
		return uuid.Nil, err
	}

	s.invalidateEngine(orgID)
	return policy.ID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelService) ListPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	policies, err := s.repo.ListPolicies(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToPolicyViews(policies), nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListOrgPolicies implements [PolicyService]: the policies registered against
// the ORGANIZATION itself — the ones that reach every employee.
//
// This is a read of the org register (org_sentinel_policies), not of the policy
// catalog. ListPolicies above returns every policy the org owns, most of which
// reach nobody or reach only a role/group/user. The two are different sets and
// the difference is the whole point of attachment-driven evaluation.
func (s *sentinelService) ListOrgPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	policies, err := s.repo.ListPoliciesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToPolicyViews(policies), nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetPolicy implements [SentinelPolicyService]. The repo scopes by org, so
// another tenant's policy comes back as not-found rather than being fetched and
// then rejected here.
func (s *sentinelService) GetPolicy(ctx context.Context, orgID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	policy, err := s.repo.GetPolicy(ctx, orgID, policyID)
	if err != nil {
		return sentinelDtos.PolicyView{}, err
	}
	return sentinelDtos.ToPolicyView(&policy), nil
}

/* ---------------------------------- <***> --------------------------------- */

// AttachPolicies implements [SentinelPolicyService]: register existing policies
// against a role, so they reach everyone holding it. Runs in one transaction so
// a partial attach never persists.
//
// Each policy is fetched org-scoped first: the register itself does not verify
// that a policy id belongs to the org, so skipping this would let a caller
// attach another tenant's policy.
func (s *sentinelService) AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	if len(policyIDs) == 0 {
		return nil
	}

	if _, err := s.repo.GetRole(ctx, orgID, roleID); err != nil {
		return fmt.Errorf("%w: %s", sentinelServiceErrs.RoleNotFoundInOrgErr, roleID)
	}

	if err := s.withTx(ctx, func(txCtx context.Context) error {
		for _, policyID := range policyIDs {
			if _, err := s.repo.GetPolicy(txCtx, orgID, policyID); err != nil {
				return fmt.Errorf("%w: %s", sentinelServiceErrs.PolicyNotFoundInOrgErr, policyID)
			}
			if err := s.repo.AttachPolicyToRole(txCtx, orgID, roleID, policyID); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	s.invalidateEngine(orgID)
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// DetachPolicy implements [SentinelPolicyService]: stop a policy reaching a
// role. The policy row itself is untouched, and the registration is soft-deleted
// so it can be restored with a history.
func (s *sentinelService) DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	if err := s.repo.DetachPolicyFromRole(ctx, orgID, roleID, policyID); err != nil {
		return err
	}
	s.invalidateEngine(orgID)
	return nil
}
