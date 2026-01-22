package hospitality_order_engine_service

import hospitality_order_engine_repository "sharpnr-engine/modules/hospitality/order_engine/repository"

type hospitalityOrderEngineService struct {
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface
}

func NewHospitalityOrderEngineService(
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface,
) *hospitalityOrderEngineService {
	return &hospitalityOrderEngineService{
		hospOrderEngRepo: hospOrderEngRepo,
	}
}
