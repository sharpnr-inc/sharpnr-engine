package kitchenservices

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
)

type kitchenService struct {
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewKitchenService(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
) *kitchenService {
	return &kitchenService{
		kitchenRepo: kitchenRepo,
	}
}
