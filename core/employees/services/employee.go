package coreEmployeeServices

import (
	"context"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"time"

	"github.com/google/uuid"
)

// CreateDefaultSuperAdmin implements HrEmployeeServiceInterface.
func (h *employeeService) CreateDefaultSuperAdmin(
	ctx context.Context,
	orgId uuid.UUID,
	rootEmail string,
	hashPass string,
) (uuid.UUID, error) {

	employee := &coreEmployeeModels.Employee{
		ID:             uuid.New(),
		Email:          rootEmail,
		OrganizationID: orgId,
		Status:         "ACTIVE",
	}

	if err := h.employeeRepo.CreateEmployee(ctx, employee); err != nil {
		return uuid.Nil, err
	}

	// 2️⃣ Create Auth (password provisioning can be deferred)
	auth := &coreEmployeeModels.EmployeeAuth{
		ID:            uuid.New(),
		EmployeeID:    employee.ID,
		Email:         rootEmail,
		Username:      rootEmail,
		IsActive:      true,
		IsProvisioned: true,
		PasswordHash:  hashPass,
	}

	if err := h.employeeRepo.CreateEmployeeAuth(ctx, auth); err != nil {
		return uuid.Nil, err
	}

	onboarding := &coreEmployeeModels.EmployeeOnboarding{
		ID:         uuid.New(),
		EmployeeID: employee.ID,
		Status:     "COMPLETED",
		CreatedAt:  time.Now(),
	}

	if err := h.employeeRepo.CreateEmployeeOnBoarding(ctx, onboarding); err != nil {
		return uuid.Nil, err
	}

	return employee.ID, nil
}

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

func buildEmployeeName(firstName, lastName, nickName string) string {
	if nickName != "" {
		return nickName
	}

	if lastName != "" {
		return firstName + " " + lastName
	}

	return firstName
}

/* ---------------------------------- <***> --------------------------------- */
