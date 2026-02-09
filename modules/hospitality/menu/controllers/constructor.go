package menu_controllers

import menuservice "sharpnr-engine/modules/hospitality/menu/service"

type menuController struct {
	menuService menuservice.MenuService
}

func NewMenuController(menuService menuservice.MenuService) *menuController {
	return &menuController{
		menuService: menuService,
	}
}
