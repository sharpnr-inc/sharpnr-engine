package hotel_sere_adapter

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"sharpnr-engine/sere"
)

func RegisterHotelWorkFlow(
	s *sere.SERE,
	hotelRepo hotelrepo.HotelRepositoryInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", defaultHotelHandler(hotelRepo))
}
