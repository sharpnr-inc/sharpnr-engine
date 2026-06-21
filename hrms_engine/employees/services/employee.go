package employee_services

import (
	"context"
	sharphr_dtos "sharpnr-engine/hrms_engine/employees/dtos"
	employee_models "sharpnr-engine/hrms_engine/employees/models"
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

	if err := h.employeeRepo.CreateEmployee(ctx, employee); err != nil {
		return uuid.Nil, err
	}

	// 2️⃣ Create Auth (password provisioning can be deferred)
	auth := &employee_models.EmployeeAuth{
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

	onboarding := &employee_models.EmployeeOnboarding{
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
) ([]*sharphr_dtos.OrgEmployeeList, error) {

	empls, err := s.employeeRepo.ListEmployeesByOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	var result []*sharphr_dtos.OrgEmployeeList

	for _, e := range empls {
		result = append(result, &sharphr_dtos.OrgEmployeeList{
			ID:           e.ID,
			EmployeeName: buildEmployeeName(e.FirstName, e.LastName, e.NickName),
			EmployeeRole: e.Role,
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
