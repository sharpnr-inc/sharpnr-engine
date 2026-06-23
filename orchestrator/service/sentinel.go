package orchestrator_service

import (
	"context"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelservice "sharpnr-engine/sentinel_engine/service"

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

func (s *orchestratorService) ListManagedPolicies(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelservice.ManagedPolicyView, error) {
	return s.sentinelService.ListManagedPolicies(ctx, orgID)
}

func (s *orchestratorService) CreateManagedPolicy(ctx context.Context, orgID, emplID uuid.UUID, in sentinelservice.CreateManagedPolicyInput) (uuid.UUID, error) {
	return s.sentinelService.CreateManagedPolicy(ctx, orgID, in)
}

func (s *orchestratorService) ListRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelservice.RoleView, error) {
	return s.sentinelService.ListRoles(ctx, orgID)
}

func (s *orchestratorService) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelservice.CreateRoleInput) (uuid.UUID, error) {
	return s.sentinelService.CreateRole(ctx, orgID, emplID, in)
}

func (s *orchestratorService) GetRolePrivileges(ctx context.Context, orgID, emplID, roleID uuid.UUID) (sentinelservice.RolePrivilegesView, error) {
	return s.sentinelService.GetRolePrivileges(ctx, orgID, roleID)
}

func (s *orchestratorService) AttachPoliciesToRole(ctx context.Context, orgID, emplID, roleID uuid.UUID, in sentinelservice.AttachPoliciesInput) error {
	return s.sentinelService.AttachPolicies(ctx, orgID, emplID, roleID, in.ManagedPolicyIDs)
}

func (s *orchestratorService) DetachPolicyFromRole(ctx context.Context, orgID, emplID, roleID, policyID uuid.UUID) error {
	return s.sentinelService.DetachPolicy(ctx, orgID, roleID, policyID)
}
