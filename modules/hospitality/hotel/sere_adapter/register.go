package hotel_sere_adapter

import (
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"sharpnr-engine/sere"
)

func RegisterHotelWorkFlow(
	s *sere.SERE,
	hotelRepo hotelrepo.HotelRepositoryInterface,
	hospitalityOrgFactoryRepository hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", defaultHotelHandler(hotelRepo))
}
