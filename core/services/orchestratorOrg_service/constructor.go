package orchestrator_org_service

import (
	"sharpnr-engine/core/interfaces"
)

type orchestratorOrgService struct {
	orgRepo interfaces.OrganizationRepositoryInterface
}

// GetOrgDetails implements interfaces.OrchestratorOrganizationServiceInterface.
func (o *orchestratorOrgService) GetOrgDetails() {
	panic("unimplemented")
}

func NewOrchestratorOrgService(orgRepo interfaces.OrganizationRepositoryInterface) *orchestratorOrgService {
	return &orchestratorOrgService{
		orgRepo: orgRepo,
	}
}
