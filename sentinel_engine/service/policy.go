package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// CreatePolicy persists a single raw Cedar policy authored in the frontend. The
// FE validates with cedar-wasm, but we revalidate here so a malformed policy can
// never reach the engine (which fails closed on a parse error). On success the
// org's compiled policy-set cache is dropped so the policy takes effect.
func (s *sentinelService) CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	text := strings.TrimSpace(in.Policy)
	if text == "" {
		return uuid.Nil, fmt.Errorf("policy text is required")
	}

	var p cedar.Policy
	if err := p.UnmarshalCedar([]byte(text)); err != nil {
		return uuid.Nil, fmt.Errorf("invalid cedar policy: %w", err)
	}

	policy := &sentinelengine_models.SentinelPolicy{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           in.Name,
		Description:    in.Description,
		Policy:         text,
		Enabled:        in.IsEnabled,
	}
	if err := s.repo.CreatePolicy(ctx, policy); err != nil {
		return uuid.Nil, fmt.Errorf("create policy: %w", err)
	}

	s.engine.InvalidateOrg(orgID)
	return policy.ID, nil
}

// ListPolicies returns every policy authored for an org (enabled or not).
func (s *sentinelService) ListPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	rows, err := s.repo.ListPoliciesByOrgID(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("list policies: %w", err)
	}
	return sentinelDtos.ToPolicyViews(rows), nil
}

// GetPolicy returns a single policy, guarding that it belongs to the org so one
// tenant can never read another's policy by id.
func (s *sentinelService) GetPolicy(ctx context.Context, orgID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	policy, err := s.repo.GetPolicyByID(ctx, policyID)
	if err != nil {
		return sentinelDtos.PolicyView{}, fmt.Errorf("get policy %s: %w", policyID, err)
	}
	if policy.OrganizationID != orgID {
		return sentinelDtos.PolicyView{}, fmt.Errorf("policy %s does not belong to org %s", policyID, orgID)
	}
	return sentinelDtos.ToPolicyView(&policy), nil
}
