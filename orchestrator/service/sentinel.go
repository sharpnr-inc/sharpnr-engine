package orchestrator_service

import (
	"context"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"

	"github.com/google/uuid"
)

// Sentinel (authorization) orchestration. These methods are thin pass-throughs
// to the Sentinel service — the orchestrator is the HTTP facade the frontend
// talks to, and authorization authoring lives in sentinel_engine. orgID/emplID
// are carried for tenant scope and future per-call authorization, even where a
// delegate doesn't need them yet.

func (s *orchestratorService) ListPermissionCatalog(ctx context.Context) []sentinel_actions.Permission {
	return s.sentinelService.Catalog()
}

func (s *orchestratorService) ListPolicies(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	return s.sentinelService.ListPolicies(ctx, orgID)
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
