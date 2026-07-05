package hrmsEmployeeServices

import (
	"context"
	hrmsEmployeesDtos "sharpnr-engine/hrms_engine/employees/dtos"

	"github.com/google/uuid"
)

// ListOrgAllEmployees implements HRMSEmployeeServiceInterface.
func (s *hrmsemployeeService) ListOrgAllEmployees(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]*hrmsEmployeesDtos.OrgEmployeeList, error) {

	empls, err := s.employeeService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}

	result := make([]*hrmsEmployeesDtos.OrgEmployeeList, 0, len(empls))

	for _, e := range empls {
		if e == nil {
			continue
		}

		result = append(result, &hrmsEmployeesDtos.OrgEmployeeList{
			ID:           e.ID,
			EmployeeName: e.EmployeeName,
			EmployeeID:   e.EmployeeID,
			Status:       e.Status,
		})
	}

	return result, nil
}
