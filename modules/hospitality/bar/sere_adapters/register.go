package bar_sere_adapters

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	"sharpnr-engine/sere"
)

func RegisterBarWorkFlow(
	s *sere.SERE,
	barRepo barrepo.BarRepository,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", defaultBarHandler(barRepo))
}
