package sentinelservice

import (
	"context"
	"github.com/google/uuid"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinel_resources "sharpnr-engine/sentinel/resources"
)

type sentinelService struct {
	repo      sentinel_repo.SentinelRepositoryInterface
	txManager dbtx.TxManager
}

// AttachPolicies implements [SentinelService].
func (s *sentinelService) AttachPolicies(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	panic("unimplemented")
}

// CreateGroup implements [SentinelService].
func (s *sentinelService) CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error) {
	panic("unimplemented")
}

// CreatePolicy implements [SentinelService].
func (s *sentinelService) CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	panic("unimplemented")
}

// CreateRole implements [SentinelService].
func (s *sentinelService) CreateRole(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	panic("unimplemented")
}

// DetachPolicy implements [SentinelService].
func (s *sentinelService) DetachPolicy(ctx context.Context, orgID uuid.UUID, roleID uuid.UUID, policyID uuid.UUID) error {
	panic("unimplemented")
}

// GetEmployeeRole implements [SentinelService].
func (s *sentinelService) GetEmployeeRole(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) (sentinelDtos.RoleView, error) {
	panic("unimplemented")
}

// GetPolicy implements [SentinelService].
func (s *sentinelService) GetPolicy(ctx context.Context, orgID uuid.UUID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	panic("unimplemented")
}

// GetPrincipals implements [SentinelService].
func (s *sentinelService) GetPrincipals(ctx context.Context, orgID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	panic("unimplemented")
}

// GetResourceActions implements [SentinelService].
func (s *sentinelService) GetResourceActions(resourceType string) ([]sentinel_resources.Action, error) {
	panic("unimplemented")
}

// GetResourceAttributes implements [SentinelService].
func (s *sentinelService) GetResourceAttributes(resourceType string) ([]sentinel_resources.Attribute, error) {
	panic("unimplemented")
}

// GetRolePrivileges implements [SentinelService].
func (s *sentinelService) GetRolePrivileges(ctx context.Context, orgID uuid.UUID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	panic("unimplemented")
}

// ListGroups implements [SentinelService].
func (s *sentinelService) ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	panic("unimplemented")
}

// ListPolicies implements [SentinelService].
func (s *sentinelService) ListPolicies(ctx context.Context, orgID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]sentinelDtos.PolicyView, error) {
	panic("unimplemented")
}

// ListResourceInstances implements [SentinelService].
func (s *sentinelService) ListResourceInstances(ctx context.Context, orgID uuid.UUID, resourceType string) ([]sentinel_resources.Ref, error) {
	panic("unimplemented")
}

// ListResources implements [SentinelService].
func (s *sentinelService) ListResources() []sentinel_resources.Summary {
	panic("unimplemented")
}

// ListRoles implements [SentinelService].
func (s *sentinelService) ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	panic("unimplemented")
}

// ProvisionOrganization implements [SentinelService].
func (s *sentinelService) ProvisionOrganization(ctx context.Context, orgID uuid.UUID, firstUserID uuid.UUID) error {
	panic("unimplemented")
}

func NewSentinelService(
	repo sentinel_repo.SentinelRepositoryInterface,
	txManager dbtx.TxManager,
) *sentinelService {
	return &sentinelService{
		repo:      repo,
		txManager: txManager,
	}
}
