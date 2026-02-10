package menuservice

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/repo"
)

type menuService struct {
	menuRepo menurepo.MenuRepositoryInterface
}

func NewHospitalityMenuService(menuRepo menurepo.MenuRepositoryInterface) *menuService {
	return &menuService{
		menuRepo: menuRepo,
	}

}
