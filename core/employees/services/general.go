package coreEmployeeServices

import (
	"context"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"

	"github.com/google/uuid"
)

// IsUserExists implements EmployeeServiceInterface.
func (h *employeeService) IsUserExists(
	ctx context.Context,
	employeeEmail string,
) (bool, error) {
	return h.employeeRepo.EmployeeExistsByEmail(ctx, employeeEmail)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *employeeService) ListOrgAllEmployees(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]*coreEmployeeContracts.OrgEmployeeList, error) {

	empls, err := s.employeeRepo.ListEmployeesByOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	var result []*coreEmployeeContracts.OrgEmployeeList

	for _, e := range empls {
		result = append(result, &coreEmployeeContracts.OrgEmployeeList{
			ID:           e.ID,
			EmployeeName: buildEmployeeName(e.FirstName, e.LastName, e.NickName),
			EmployeeID:   e.EmployeeCode,
			Status:       e.Status,
		})
	}

	return result, nil
}
