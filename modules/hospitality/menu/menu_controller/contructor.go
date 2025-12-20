package menucontroller

import menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"

type menuController struct {
	menuService menuservice.MenuService
}

func NewMenuController(menuService menuservice.MenuService) *menuController {
	return &menuController{
		menuService: menuService,
	}
}
