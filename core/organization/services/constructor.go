package organization_services

import (
	"context"
	organization_repo "sharpnr-engine/core/organization/repo"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationService struct {
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo  organization_repo.OrganizationRepositoryInterface
}

// CreateOrganization implements OrganizationInternalServiceInterface.
func (r *organizationService) CreateOrganization(ctx context.Context, input interface{}) error {
	panic("unimplemented")
}

// GetOrgHQ implements OrganizationPublicServiceInterface.
func (r *organizationService) GetOrgHQ() {
	panic("unimplemented")
}

func NewOrganizationService(
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface,
	orgMainRepo organization_repo.OrganizationRepositoryInterface,
) *organizationService {
	return &organizationService{
		mailerEngine: mailerEngine,
		orgRedisRepo: orgRedisRepo,
		orgMainRepo:  orgMainRepo,
	}
}
