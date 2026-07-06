package coreEmployeeServices

import (
	"context"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeModels "sharpnr-engine/core/employees/models"
	"time"

	"github.com/google/uuid"
)

func (s *employeeService) UpdateSuperAdminEmployee(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	updates *coreEmployeeContracts.UpdateSuperAdminEmployeeInput) error {
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

/* ---------------------------------- <***> --------------------------------- */

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
