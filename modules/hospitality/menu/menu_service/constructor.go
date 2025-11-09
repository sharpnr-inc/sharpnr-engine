package menuservice

import menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"

type menuService struct {
	menuRepo menurepo.MenuRepository
}

func NewHospitalityMenuService(menuRepo menurepo.MenuRepository) *menuService {
	return &menuService{
		menuRepo: menuRepo,
	}

}
