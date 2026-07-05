package orchestrator_service

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// Sentinel (authorization) orchestration. These methods are thin pass-throughs
// to the Sentinel service — the orchestrator is the HTTP facade the frontend
// talks to, and authorization authoring lives in sentinel_engine. orgID/emplID
// are carried for tenant scope and future per-call authorization, even where a
// delegate doesn't need them yet.

func (s *orchestratorService) ListPrincipals(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	return s.sentinelService.GetPrincipals(ctx, orgID)
}

func (s *orchestratorService) ListAuthzResources(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorService.ResourceSummaryView, error) {
	return s.sentinelAggregator.ListResources(ctx, orgID, emplID)
}

func (s *orchestratorService) GetAuthzResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (sentinelAggregatorService.ResourceDetailsView, error) {
	return s.sentinelAggregator.GetResourceDetails(ctx, orgID, emplID, resourceType)
}

func (s *orchestratorService) GetAuthzContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error) {
	return s.sentinelAggregator.GetContextPalette(ctx, orgID, emplID)
}

func (s *orchestratorService) ListPolicies(ctx context.Context, orgID, emplID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]sentinelDtos.PolicyView, error) {
	return s.sentinelService.ListPolicies(ctx, orgID, scope)
}

func (s *orchestratorService) CreatePolicy(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	return s.sentinelService.CreatePolicy(ctx, orgID, in)
}

func (s *orchestratorService) ListRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	return s.sentinelService.ListRoles(ctx, orgID)
}

func (s *orchestratorService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	return s.sentinelService.CreateRole(ctx, orgID, emplID, in)
}

func (s *orchestratorService) GetRolePrivileges(ctx context.Context, orgID, emplID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	return s.sentinelService.GetRolePrivileges(ctx, orgID, roleID)
}

func (s *orchestratorService) AttachPoliciesToRole(ctx context.Context, orgID, emplID, roleID uuid.UUID, in sentinelDtos.AttachPoliciesInput) error {
	return s.sentinelService.AttachPolicies(ctx, orgID, emplID, roleID, in.PolicyIDs)
}

func (s *orchestratorService) DetachPolicyFromRole(ctx context.Context, orgID, emplID, roleID, policyID uuid.UUID) error {
	return s.sentinelService.DetachPolicy(ctx, orgID, roleID, policyID)
}
