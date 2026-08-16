package hrmsEngineService

import (
	"context"
	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"

	"github.com/google/uuid"
)

// ListOrgAllEmployees implements HrmsEngineServiceInterface.
func (s *hrmsService) ListOrgAllEmployees(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]*hrmsEngineDtos.OrgEmployeeList, error) {

	empls, err := s.employeeService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}

	result := make([]*hrmsEngineDtos.OrgEmployeeList, 0, len(empls))

	for _, e := range empls {
		if e == nil {
			continue
		}

		result = append(result, &hrmsEngineDtos.OrgEmployeeList{
			ID:           e.ID,
			EmployeeName: e.EmployeeName,
			EmployeeID:   e.EmployeeID,
			Status:       e.Status,
		})
	}

	return result, nil
}
