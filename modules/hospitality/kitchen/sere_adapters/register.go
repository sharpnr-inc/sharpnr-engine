package kitchen_sere_adapters

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
)

func RegisterKitchenWorkFlow(s *sere.SERE, kitchenRepo kitchenrepo.KitchenRepositoryInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", createDefaultKitchen(kitchenRepo))
}
