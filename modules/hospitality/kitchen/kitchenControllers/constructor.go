package kitchencontrollers

import kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"

type kitchenController struct {
	kitchenService kitchenservices.KitchenServiceInterface
}

func NewKitchenController(kitchenService kitchenservices.KitchenServiceInterface) *kitchenController {
	return &kitchenController{
		kitchenService: kitchenService,
	}
}
