package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

// Fetches all the employee of the organization
// Fetches all the employees of the organization
func (s *orchestratorService) GetAllEmployeesOfOrg(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]orchestrator_dtos.EmployeeOfOrganization, error) {

	// this will return all the employees in the organization
	// NOTE: we have to wire the sentinal engine here for the authorization which we will do it later

	empList, err := s.employeeHRService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}

	// Pre-allocate the slice to optimize memory and performance
	result := make([]orchestrator_dtos.EmployeeOfOrganization, 0, len(empList))

	for _, emp := range empList {
		// Safety check to avoid nil pointer dereference
		if emp == nil {
			continue
		}

		// Map the fields from OrgEmployeeList to EmployeeOfOrganization
		mappedEmp := orchestrator_dtos.EmployeeOfOrganization{
			ID:   emp.ID,
			Name: emp.EmployeeName,
			// Nickname is omitted as it doesn't exist in OrgEmployeeList (defaults to "")
			Role: orchestrator_dtos.EmployeeRole{
				Role: emp.EmployeeRole,
			},
			Status: emp.Status,
		}

		result = append(result, mappedEmp)
	}

	return result, nil
}
