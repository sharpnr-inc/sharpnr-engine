package kitchencontrollers

import (
	kitchenservices "sharpnr-engine/business_ops/hospitality/kitchen/services"
	sse_engine "sharpnr-engine/sse_engine"
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
