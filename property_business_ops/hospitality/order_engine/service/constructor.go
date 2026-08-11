package hospitality_order_engine_service

import (
	"context"
	kitchenservices "sharpnr-engine/property_infra/hospitality_infra/kitchen/services"
	menuservice "sharpnr-engine/property_business_ops/hospitality/menu/service"
	hospitality_order_engine_contracts "sharpnr-engine/property_business_ops/hospitality/order_engine/contracts"
	hospitality_order_engine_repository "sharpnr-engine/property_business_ops/hospitality/order_engine/repository"
	"sharpnr-engine/sero"
)

type hospitalityOrderEngineService struct {
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface
	seroEngine       *sero.SEROEngine
	menuService      menuservice.MenuService
	kitchenService   kitchenservices.KitchenServiceInterface
}

func NewHospitalityOrderEngineService(
	hospOrderEngRepo hospitality_order_engine_repository.HospitalityOrderEngineRepositoryInterface,
	seroEngine *sero.SEROEngine,
	menuService menuservice.MenuService,
	kitchenService kitchenservices.KitchenServiceInterface,
) *hospitalityOrderEngineService {
	return &hospitalityOrderEngineService{
		hospOrderEngRepo: hospOrderEngRepo,
		seroEngine:       seroEngine,
		menuService:      menuService,
		kitchenService:   kitchenService,
	}
}

// CreateOrder implements HospitalityOrderEngineServiceInterface.
func (h *hospitalityOrderEngineService) CreateOrder(ctx context.Context, input *hospitality_order_engine_contracts.CreateOrderInput) error {
	panic("unimplemented")
}
