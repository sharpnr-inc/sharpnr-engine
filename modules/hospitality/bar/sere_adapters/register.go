package bar_sere_adapters

import (
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	"sharpnr-engine/sere"
)

func RegisterBarWorkFlow(
	s *sere.SERE,
	barRepo barrepo.BarRepository,
	hospitalityOrgFactoryRepository hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", defaultBarHandler(barRepo, hospitalityOrgFactoryRepository))
}
