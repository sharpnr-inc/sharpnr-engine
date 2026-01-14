package menu_sere_adapter

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/sere"
)

func RegisterHospitalityBarMenuWorkflow(
	s *sere.SERE,
	menuRepo menurepo.MenuRepositoryInterface,
) {
}

func RegisterHospitalityRestaurantMenuWorkflow(
	s *sere.SERE,
	menuRepo menurepo.MenuRepositoryInterface,
) {
}

func RegisterHospitalityHotelMenuWorkflow(
	s *sere.SERE,
	menuRepo menurepo.MenuRepositoryInterface,
) {
}
