package organization_services

import (
	"context"
	organization_repo "sharpnr-engine/core/organization/repo"
)

type organizationService struct {
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo  organization_repo.OrganizationRepositoryInterface
}

// AddHQ implements OrganizationInternalServiceInterface.
func (r *organizationService) AddHQ(ctx context.Context, input interface{}) error {
	panic("unimplemented")
}

// CreateOrganization implements OrganizationInternalServiceInterface.
func (r *organizationService) CreateOrganization(ctx context.Context, input interface{}) error {
	panic("unimplemented")
}

// GetOrgHQ implements OrganizationPublicServiceInterface.
func (r *organizationService) GetOrgHQ() {
	panic("unimplemented")
}

// IsOrgExists implements OrganizationPublicServiceInterface.
func (r *organizationService) IsOrgExists(ctx context.Context, rootEmail string) (bool, error) {
	panic("unimplemented")
}

func NewOrganizationService(
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface,
	orgMainRepo organization_repo.OrganizationRepositoryInterface,
) *organizationService {
	return &organizationService{
		orgRedisRepo: orgRedisRepo,
		orgMainRepo:  orgMainRepo,
	}
}
