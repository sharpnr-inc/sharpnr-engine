package kitchencontrollers

import (
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	sse_engine "sharpnr-engine/sse"
)

type kitchenController struct {
	kitchenService kitchenservices.KitchenServiceInterface
	sseEngine      *sse_engine.StreamManager
}

func NewKitchenController(kitchenService kitchenservices.KitchenServiceInterface, sseEngine *sse_engine.StreamManager) *kitchenController {
	return &kitchenController{
		kitchenService: kitchenService,
		sseEngine:      sseEngine,
	}
}
