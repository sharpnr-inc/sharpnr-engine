package restaurant_sere_adapters

import (
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/sere"
)

func RegisterRestaurantWorkFlow(
	s *sere.SERE,
	restaurantRepo restaurantrepo.RestaurantRepository,
	hospitalityOrgFactoryRepository hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", createDefaultRestaurant(restaurantRepo, hospitalityOrgFactoryRepository))
}
