package employee_services

import (
	employee_repo "sharpnr-engine/modules/hr/employees/repo"
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

// NewHrEmployeeService constructs a new employeeService.
//
// All repositories are required.
// If you don’t need one yet, pass nil ONLY if the method
// using it is not called (documented per method).
func NewEmployeeService(
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
