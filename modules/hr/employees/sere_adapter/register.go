package employee_sere_adapter

import (
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
)

func RegisterHREmployeeWorkflows(s *sere.SERE, hrService hr_employee_services.EmployeeServiceInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", hrEmployeeOnboardHandler(hrService))
}
