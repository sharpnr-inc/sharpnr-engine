package hrmsEmployeeServices

import (
	"context"
	hrmsEmployeesDtos "sharpnr-engine/hrms_engine/employees/dtos"

	"github.com/google/uuid"
)

type HRMSEmployeeServiceInterface interface {

	// Here we will define the methods that will be implemented by the HRMSEmployeeService struct. These methods will be related to employee management in the HRMS context.
	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*hrmsEmployeesDtos.OrgEmployeeList, error)
}
