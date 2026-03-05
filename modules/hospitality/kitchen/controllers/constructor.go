package kitchencontrollers

import (
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	sse_engine "sharpnr-engine/sse"
)

type KitchenController struct {
	kitchenService kitchenservices.KitchenServiceInterface
	sseEngine      *sse_engine.StreamManager
}

func NewKitchenController(kitchenService kitchenservices.KitchenServiceInterface, sseEngine *sse_engine.StreamManager) *KitchenController {
	return &KitchenController{
		kitchenService: kitchenService,
		sseEngine:      sseEngine,
	}
}
