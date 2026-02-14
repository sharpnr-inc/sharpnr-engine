package hospitality_order_engine_service

import (
	"context"
	hospitality_order_engine_contracts "sharpnr-engine/modules/hospitality/order_engine/contracts"
	hospitality_order_engine_repository "sharpnr-engine/modules/hospitality/order_engine/repository"
	"sharpnr-engine/sero"
)

type hospitalityOrderEngineService struct {
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface
	seroEngine       *sero.SEROEngine
}

// CreateOrder implements HospitalityOrderEngineServiceInterface.
func (h *hospitalityOrderEngineService) CreateOrder(ctx context.Context, input *hospitality_order_engine_contracts.CreateOrderInput) error {
	panic("unimplemented")
}

func NewHospitalityOrderEngineService(
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface,
	seroEngine *sero.SEROEngine,
) *hospitalityOrderEngineService {
	return &hospitalityOrderEngineService{
		hospOrderEngRepo: hospOrderEngRepo,
		seroEngine:       seroEngine,
	}
}
