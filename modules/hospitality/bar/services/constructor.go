package barservices

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
)

type barServices struct {
	barRepo     barrepo.BarRepository
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewBarServices(
	barRepo barrepo.BarRepository,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface) *barServices {
	return &barServices{
		barRepo:     barRepo,
		kitchenRepo: kitchenRepo,
	}
}
