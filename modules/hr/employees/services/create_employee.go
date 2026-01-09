package hr_employee_services

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"time"

	"github.com/google/uuid"
)

// CreateDefaultSuperAdmin implements HrEmployeeServiceInterface.
func (h *hrEmployeeService) CreateDefaultSuperAdmin(
	ctx context.Context,
	orgId uuid.UUID,
	rootEmail string,
	hashPass string,
) (uuid.UUID, error) {

	employee := &sharphr_models.Employee{
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
	auth := &sharphr_models.EmployeeAuth{
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

	onboarding := &sharphr_models.EmployeeOnboarding{
		ID:         uuid.New(),
		EmployeeID: employee.ID,
		Status:     "COMPLETED",
		CreatedAt:  time.Now(),
	}

	if err := h.empOnboardingRepo.Create(ctx, onboarding); err != nil {
		return uuid.Nil, err
	}

	// 4️⃣ Create OWNER policy
	policy := &sharphr_models.EmployeePolicy{
		ID:             uuid.New(),
		OrganizationID: orgId,
		Code:           "OWNER",
		Name:           "Organization Owner",
		Type:           "SYSTEM",
	}

	if err := h.policyRepo.Create(ctx, policy); err != nil {
		return uuid.Nil, err
	}

	// // 6️⃣ Assign policy to employee
	// assignment := &sharphr_models.EmployeePrivilegePolicy{
	// 	ID:               uuid.New(),
	// 	EmployeeID:       employee.ID,
	// 	EmployeePolicyID: policy.ID,
	// 	OrganizationID:   orgId,
	// 	Domain:           "HOSPITALITY",
	// 	AssignedAt:       time.Now(),
	// }

	return employee.ID, nil
}
