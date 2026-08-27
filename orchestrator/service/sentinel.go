package orchestrator_service

import (
	"context"

	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// Sentinel (authorization) orchestration. These methods are thin pass-throughs
// to the Sentinel service — the orchestrator is the HTTP facade the frontend
// talks to, and authorization authoring lives in sentinel_engine. orgID/emplID
// are carried for tenant scope and future per-call authorization, even where a
// delegate doesn't need them yet.

// sentinelOrchestrator implements SentinelOrchestrationInterface.
//
// guard is the seam for the missing authoring authorization: policy/role
// creation and attachment must call guard.Authorize before delegating (rule 1
// in doc.go). The checks are not wired yet — do not ship policy authoring to
// production before they are.
type sentinelOrchestrator struct {
	guard              *orchestrator_authz.Guard
	sentinelService    sentinelservice.SentinelServiceInterface
	sentinelAggregator sentinelAggregatorService.SentinelAggregatorServiceInterface
}

func (s *sentinelOrchestrator) ListPrincipals(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	return s.sentinelService.GetPrincipals(ctx, orgID)
}

func (s *sentinelOrchestrator) ListAuthzResources(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorService.ResourceSummaryView, error) {
	return s.sentinelAggregator.ListResources(ctx, orgID, emplID)
}

func (s *sentinelOrchestrator) GetAuthzResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (sentinelAggregatorService.ResourceDetailsView, error) {
	return s.sentinelAggregator.GetResourceDetails(ctx, orgID, emplID, resourceType)
}

func (s *sentinelOrchestrator) GetAuthzContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error) {
	return s.sentinelAggregator.GetContextPalette(ctx, orgID, emplID)
}

func (s *sentinelOrchestrator) CreatePolicy(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	return s.sentinelService.CreatePolicy(ctx, orgID, in)
}

func (s *sentinelOrchestrator) ListPolicies(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	return s.sentinelService.ListPolicies(ctx, orgID)
}

func (s *sentinelOrchestrator) GetPolicy(ctx context.Context, orgID, emplID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	return s.sentinelService.GetPolicy(ctx, orgID, policyID)
}

func (s *sentinelOrchestrator) ListRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	return s.sentinelService.ListRoles(ctx, orgID)
}

func (s *sentinelOrchestrator) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	return s.sentinelService.CreateRole(ctx, orgID, emplID, in)
}

func (s *sentinelOrchestrator) GetRolePrivileges(ctx context.Context, orgID, emplID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	return s.sentinelService.GetRolePrivileges(ctx, orgID, roleID)
}

func (s *sentinelOrchestrator) AttachPoliciesToRole(ctx context.Context, orgID, emplID, roleID uuid.UUID, in sentinelDtos.AttachPoliciesInput) error {
	return s.sentinelService.AttachPolicies(ctx, orgID, emplID, roleID, in.PolicyIDs)
}

func (s *sentinelOrchestrator) DetachPolicyFromRole(ctx context.Context, orgID, emplID, roleID, policyID uuid.UUID) error {
	return s.sentinelService.DetachPolicy(ctx, orgID, roleID, policyID)
}
