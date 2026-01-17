package organization_services

import organization_repo "sharpnr-engine/core/organization/repo"

type organizationService struct {
	orgTempRepo organization_repo.TempOrgSignUpRepositoryInterface
	orgMainRepo organization_repo.MainOrgSignUpRepositoryInterface
}

func NewOrganizationService(
	orgTempRepo organization_repo.TempOrgSignUpRepositoryInterface,
	orgMainRepo organization_repo.MainOrgSignUpRepositoryInterface,
) *organizationService {
	return &organizationService{
		orgTempRepo: orgTempRepo,
		orgMainRepo: orgMainRepo,
	}
}
