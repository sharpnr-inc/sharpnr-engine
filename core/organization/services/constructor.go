package organization_services

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	sharpIam_services "sharpnr-engine/core/sharp_iam/services"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationService struct {
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo  organization_repo.OrganizationRepositoryInterface
	iam          *sharpIam_services.SharpIamService
}

// GetOrgHQ implements OrganizationPublicServiceInterface.
func (r *organizationService) GetOrgHQ() {
	panic("unimplemented")
}

func NewOrganizationService(
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface,
	orgMainRepo organization_repo.OrganizationRepositoryInterface,
	iam *sharpIam_services.SharpIamService,
) *organizationService {
	return &organizationService{
		mailerEngine: mailerEngine,
		orgRedisRepo: orgRedisRepo,
		orgMainRepo:  orgMainRepo,
		iam:          iam,
	}
}
