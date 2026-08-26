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

// ListPolicies implements [SentinelPolicyService].
func (s *sentinelService) ListPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error) {

	policies, err := s.repo.ListPoliciesByOrganizationID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToPolicyViews(policies), nil
}

// GetPolicy implements [SentinelPolicyService].
func (s *sentinelService) GetPolicy(ctx context.Context, orgID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	policy, err := s.repo.GetPolicyByID(ctx, policyID)
	if err != nil {
		return sentinelDtos.PolicyView{}, err
	}
	if policy.OrganizationID != orgID {
		return sentinelDtos.PolicyView{}, sentinelServiceErrs.PolicyNotFoundInOrgErr
	}
	return sentinelDtos.ToPolicyView(&policy), nil
}

// AttachPolicies implements [SentinelPolicyService]: group existing policies
// under a role. Every policy must belong to the org; runs in one transaction so
// a partial attach never persists.
func (s *sentinelService) AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	if len(policyIDs) == 0 {
		return nil
	}

	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return err
	}
	if role.OrganizationID != orgID {
		return sentinelServiceErrs.RoleNotFoundInOrgErr
	}

	if err := s.withTx(ctx, func(txCtx context.Context) error {
		for _, policyID := range policyIDs {
			policy, err := s.repo.GetPolicyByID(txCtx, policyID)
			if err != nil {
				return fmt.Errorf("policy %s: %w", policyID, err)
			}
			if policy.OrganizationID != orgID {
				return fmt.Errorf("%w: %s", sentinelServiceErrs.PolicyNotFoundInOrgErr, policyID)
			}
			edge := &sentinelengine_models.RolePolicies{
				OrganizationID:   orgID,
				RoleID:           roleID,
				SentinelPolicyID: policyID,
			}
			if err := s.repo.AssignPolicyToRole(txCtx, edge); err != nil {
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

// DetachPolicy implements [SentinelPolicyService]: remove one role↔policy
// grouping, leaving the policy row untouched.
func (s *sentinelService) DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	if err := s.repo.RemovePolicyFromRole(ctx, orgID, roleID, policyID); err != nil {
		return err
	}
	s.invalidateEngine(orgID)
	return nil
}
