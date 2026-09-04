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

/* ---------------------------------- <***> --------------------------------- */

// ListOrgEmployees implements EmployeeIamEngineServiceInterface.
//
// The plain-identity twin of ListOrgAllEmployees: same rows, no per-employee
// role lookup. That difference matters — the decorated version issues one
// authorization query per employee and errors outright when someone has no
// role, neither of which a consumer that just wants names should inherit.
func (s *employeeService) ListOrgEmployees(
	ctx context.Context,
	orgID uuid.UUID,
) ([]coreEmployeeContracts.OrgEmployeeIdentity, error) {

	employees, err := s.employeeRepo.ListEmployeesByOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	result := make([]coreEmployeeContracts.OrgEmployeeIdentity, 0, len(employees))

	for _, e := range employees {
		if e == nil {
			continue
		}

		result = append(result, coreEmployeeContracts.OrgEmployeeIdentity{
			ID:           e.ID,
			OrgID:        e.OrganizationID,
			DisplayName:  buildEmployeeName(e.FirstName, e.LastName, e.NickName),
			Email:        e.Email,
			Status:       string(e.Status),
			EmployeeCode: e.EmployeeCode,
		})
	}

	return result, nil
}
