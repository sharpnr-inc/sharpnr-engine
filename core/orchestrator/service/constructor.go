package orchestrator_service

import hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"

type orchestratorService struct {
	hospitality_HotelRepo hotelrepo.HotelRepositoryInterface
}

func NewSharpnrOrchestratorService(
	hospitality_HotelRepo hotelrepo.HotelRepositoryInterface,
) *orchestratorService {
	return &orchestratorService{
		hospitality_HotelRepo: hospitality_HotelRepo,
	}
}
