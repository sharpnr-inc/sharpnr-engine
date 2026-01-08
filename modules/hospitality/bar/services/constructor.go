package barservices

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
)

type barServices struct {
	barRepo barrepo.BarRepository
}

func NewBarServices(
	barRepo barrepo.BarRepository) *barServices {
	return &barServices{
		barRepo: barRepo,
	}
}
