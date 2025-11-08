package barcontrollers

import barservices "sharpnr-engine/modules/hospitality/bar/services"

type barController struct {
	barService barservices.BarServices
}

func NewBarController(barService barservices.BarServices) *barController {
	return &barController{
		barService: barService,
	}
}
