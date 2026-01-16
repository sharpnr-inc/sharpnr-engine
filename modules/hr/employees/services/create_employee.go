package hr_employee_services

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"
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

	employee := &employee_models.Employee{
		ID:             uuid.New(),
		Email:          rootEmail,
		OrganizationID: orgId,
		Role:           "SUPER_ADMIN",
		Status:         "ACTIVE",
	}

	if err := h.empIdentityRepo.Create(ctx, employee); err != nil {
		return uuid.Nil, err
	}

	// 2️⃣ Create Auth (password provisioning can be deferred)
	auth := &employee_models.EmployeeAuth{
		ID:                 uuid.New(),
		EmployeeID:         employee.ID,
		Email:              rootEmail,
		Username:           rootEmail,
		IsActive:           true,
		IsProvisioned:      false,
		CanAccessDashboard: true,
		PasswordHash:       hashPass,
		CanAccessMobile:    true,
		CanAccessPOS:       true,
	}

	if err := h.empAuthRepo.Create(ctx, auth); err != nil {
		return uuid.Nil, err
	}

	onboarding := &employee_models.EmployeeOnboarding{
		ID:         uuid.New(),
		EmployeeID: employee.ID,
		Status:     "COMPLETED",
		CreatedAt:  time.Now(),
	}

	if err := h.empOnboardingRepo.Create(ctx, onboarding); err != nil {
		return uuid.Nil, err
	}

	// 4️⃣ Create OWNER policy
	policy := &employee_models.EmployeePolicy{
		ID:             uuid.New(),
		OrganizationID: orgId,
		Code:           "OWNER",
		Name:           "Organization Owner",
		Type:           "SYSTEM",
	}

	if err := h.policyRepo.Create(ctx, policy); err != nil {
		return uuid.Nil, err
	}

	return employee.ID, nil
}
