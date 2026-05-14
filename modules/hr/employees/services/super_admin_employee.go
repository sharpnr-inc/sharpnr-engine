package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"

	"github.com/google/uuid"
)

func (s *employeeService) UpdateSuperAdminEmployee(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	updates *employee_contracts.UpdateSuperAdminEmployeeInput) error {
	// 1. Convert struct pointers to a map
	updateMap := make(map[string]interface{})

	if updates.FirstName != nil {
		updateMap["first_name"] = *updates.FirstName
	}
	if updates.LastName != nil {
		updateMap["last_name"] = *updates.LastName
	}

	// Note: Usually, you wouldn't allow updating OrganizationID or EmployeeID
	// through a general update map for security, but if required:
	if updates.OrganizationID != uuid.Nil {
		updateMap["organization_id"] = updates.OrganizationID
	}

	// 2. Defensive check: If nothing is provided to update, return early
	if len(updateMap) == 0 {
		return nil
	}

	// 3. Call the repo method
	return s.employeeRepo.UpdateEmployee(ctx, orgID, emplID, updateMap)
}
