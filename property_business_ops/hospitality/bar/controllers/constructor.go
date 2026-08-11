package barcontrollers

import barservices "sharpnr-engine/property_business_ops/hospitality/bar/services"

type barController struct {
	barService barservices.BarServicesInterface
}

func NewBarController(barService barservices.BarServicesInterface) *barController {
	return &barController{
		barService: barService,
	}
}
