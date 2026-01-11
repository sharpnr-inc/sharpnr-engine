package menu_sere_adapter

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/sere"
)

func RegisterHospitalityBarMenuWorkflow(s *sere.SERE, menuRepo menurepo.MenuRepositoryInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", barDefaultMenuSetup(menuRepo))
}

func RegisterHospitalityRestaurantMenuWorkflow(s *sere.SERE, menuRepo menurepo.MenuRepositoryInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", restaurantDefaultMenuSetup(menuRepo))
}

func RegisterHospitalityHotelMenuWorkflow(s *sere.SERE, menuRepo menurepo.MenuRepositoryInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", hotelDefaultMenuSetup(menuRepo))
}
