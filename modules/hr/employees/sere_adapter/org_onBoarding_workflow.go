package employee_sere_adapter

import (
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"

	"github.com/google/uuid"
)

func RegisterHRWorkflows(s *sere.SERE, hrService hr_employee_services.HrEmployeeServiceInterface) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", hrOnboardHandler(hrService))
}

func hrOnboardHandler(
	hrService hr_employee_services.HrEmployeeServiceInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {
		orgID := wf.Data["org_id"].(uuid.UUID)
		orgRootEmail := wf.Data["email"].(string)

		empID, err := hrService.CreateDefaultSuperAdmin(wf.Ctx, orgID, orgRootEmail)
		if err != nil {
			return err
		}

		wf.Data["employee_id"] = empID
		return nil
	}
}
