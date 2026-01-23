package organization_services

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationService struct {
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo  organization_repo.OrganizationRepositoryInterface
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
