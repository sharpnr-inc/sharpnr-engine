package kitchen_sere_adapters

import (
	"sharpnr-engine/sere"
)

func RegisterKitchenWorkFlow(s *sere.SERE) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", createDefaultKitchen())
}
