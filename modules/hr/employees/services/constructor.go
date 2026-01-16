package hr_employee_services

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"
	employee_repo "sharpnr-engine/modules/hr/employees/repo"

	"github.com/google/uuid"
)

// employeeService is the orchestration layer for
// employee lifecycle operations (identity + HR concerns).
//
// IMPORTANT:
// - This service does NOT do persistence itself
// - It coordinates multiple repositories
// - Business logic lives here, not in repos
type employeeService struct {
	// Identity & Auth
	empIdentityRepo employee_repo.EmployeeIdentityRepositoryInterface
	empAuthRepo     employee_repo.EmployeeAuthRepositoryInterface

	// Onboarding & Verification
	empOnboardingRepo employee_repo.EmployeeOnboardingRepositoryInterface
	empVerifyRepo     employee_repo.EmployeeVerificationRepositoryInterface

	// HR Details
	empProfileRepo    employee_repo.EmployeeProfileRepositoryInterface
	empEmploymentRepo employee_repo.EmployeeEmploymentRepositoryInterface

	// Policy
	policyRepo           employee_repo.EmployeePolicyRepositoryInterface
	policyRuleRepo       employee_repo.EmployeePolicyRuleRepositoryInterface
	policyAssignmentRepo employee_repo.EmployeePrivilegePolicyRepositoryInterface
}

// AddEmployeeInOrg implements HrEmployeeServiceInterface.
func (h *employeeService) AddEmployeeInOrg(ctx context.Context, sharpnr_id string) error {
	panic("unimplemented")
}

// CreateNewEmployee implements HrEmployeeServiceInterface.
func (h *employeeService) CreateNewEmployee(ctx context.Context, orgID uuid.UUID) error {
	panic("unimplemented")
}

// RemoveEmployeeInOrg implements HrEmployeeServiceInterface.
func (h *employeeService) RemoveEmployeeInOrg(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) error {
	panic("unimplemented")
}

// UpdateEmployeeDetails implements HrEmployeeServiceInterface.
func (h *employeeService) UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *employee_models.Employee) (*employee_models.Employee, error) {
	panic("unimplemented")
}

// NewHrEmployeeService constructs a new employeeService.
//
// All repositories are required.
// If you don’t need one yet, pass nil ONLY if the method
// using it is not called (documented per method).
func NewHrEmployeeService(
	empIdentityRepo employee_repo.EmployeeIdentityRepositoryInterface,
	empAuthRepo employee_repo.EmployeeAuthRepositoryInterface,
	empOnboardingRepo employee_repo.EmployeeOnboardingRepositoryInterface,
	empVerifyRepo employee_repo.EmployeeVerificationRepositoryInterface,
	empProfileRepo employee_repo.EmployeeProfileRepositoryInterface,
	empEmploymentRepo employee_repo.EmployeeEmploymentRepositoryInterface,

	// Policy
	policyRepo employee_repo.EmployeePolicyRepositoryInterface,
	policyRuleRepo employee_repo.EmployeePolicyRuleRepositoryInterface,
	policyAssignmentRepo employee_repo.EmployeePrivilegePolicyRepositoryInterface,
) *employeeService {

	return &employeeService{
		empIdentityRepo:      empIdentityRepo,
		empAuthRepo:          empAuthRepo,
		empOnboardingRepo:    empOnboardingRepo,
		empVerifyRepo:        empVerifyRepo,
		empProfileRepo:       empProfileRepo,
		empEmploymentRepo:    empEmploymentRepo,
		policyRepo:           policyRepo,
		policyRuleRepo:       policyRuleRepo,
		policyAssignmentRepo: policyAssignmentRepo,
	}
}
