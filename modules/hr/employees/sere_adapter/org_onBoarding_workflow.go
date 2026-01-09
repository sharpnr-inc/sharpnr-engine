package employee_sere_adapter

import (
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"

	"github.com/google/uuid"
)

// This is the function for creating
func hrEmployeeOnboardHandler(
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
