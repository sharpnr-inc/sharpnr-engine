package kitchencontrollers

import kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"

type kitchenController struct {
	kitchenService kitchenservices.KitchenServiceInterface
}

func NewKitchenController(kitchenService kitchenservices.KitchenServiceInterface) *kitchenController {
	return &kitchenController{
		kitchenService: kitchenService,
	}
}
