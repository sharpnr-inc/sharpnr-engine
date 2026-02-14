package kitchenservices

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/repos"
	"sharpnr-engine/sero"
)

type kitchenService struct {
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
	seroEngine  *sero.SEROEngine
}

func NewKitchenService(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	seroEngine *sero.SEROEngine,
) *kitchenService {
	return &kitchenService{
		kitchenRepo: kitchenRepo,
		seroEngine:  seroEngine,
	}
}
