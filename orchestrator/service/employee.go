package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

// Fetches all the employee of the organization
func (s *orchestratorService) GetAllEmployeesOfOrg(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
) ([]orchestrator_dtos.EmployeeOfOrganization, error) {

	// this will return all the employee is the organization
	// NOTE: we have to wire the sentinal engine here for the authorization which we will do it later

	// empList, err := s.employeeHRService.ListOrgAllEmployees(ctx, orgID, emplID)
	// if err != nil {
	// 	return nil, err
	// }

	panic("")

}
