package employee_sere_adapter

import (
	"sharpnr-engine/sere"
)

func RegisterHREmployeeWorkflows(s *sere.SERE) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", hrEmployeeOnboardHandler())
}
