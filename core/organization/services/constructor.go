package organization_services

import organization_repo "sharpnr-engine/core/organization/repo"

type organizationService struct {
	orgTempRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo organization_repo.OrganizationRepositoryInterface
}

func NewOrganizationService(
	orgTempRepo organization_repo.RedisOrganizationRepositoryInterface,
	orgMainRepo organization_repo.OrganizationRepositoryInterface,
) *organizationService {
	return &organizationService{
		orgTempRepo: orgTempRepo,
		orgMainRepo: orgMainRepo,
	}
}
