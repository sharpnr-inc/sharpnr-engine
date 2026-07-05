package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

// Fetches all the employees of the organization
func (s *orchestratorService) GetAllEmployeesOfOrg(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]orchestrator_dtos.EmployeeOfOrganization, error) {

	// Fetch all the employees in the organization
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

		// Fetch the role view for this specific employee
		roleView, err := s.sentinelService.GetEmployeeRole(ctx, orgID, emp.ID)
		if err != nil {
			return nil, err
		}

		// Map the fields from OrgEmployeeList and RoleView to EmployeeOfOrganization
		mappedEmp := orchestrator_dtos.EmployeeOfOrganization{
			ID:     emp.ID,
			Name:   emp.EmployeeName,
			Status: emp.Status,
			Role: orchestrator_dtos.EmployeeRole{
				ID:          roleView.ID,
				Role:        roleView.Name, // Maps the Name from RoleView to the Role field
				Description: roleView.Description,
			},
		}

		result = append(result, mappedEmp)
	}

	return result, nil
}
