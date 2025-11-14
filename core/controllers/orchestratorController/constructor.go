package orchestratorcontroller

import "sharpnr-engine/core/interfaces"

type orchestratorOrganizationController struct {
	orchOrgSvc interfaces.OrchestratorOrganizationServiceInterface
}

func NewOrchestratorOrganizationController(orchOrgSvc interfaces.OrchestratorOrganizationServiceInterface) *orchestratorOrganizationController {
	return &orchestratorOrganizationController{
		orchOrgSvc: orchOrgSvc,
	}
}
