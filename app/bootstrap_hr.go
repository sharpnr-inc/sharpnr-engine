package app

import (
	employee_repo "sharpnr-engine/modules/hr/employees/repo"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"

	"gorm.io/gorm"
)

type hrBootstrap struct {
	Service hr_employee_services.EmployeeServiceInterface

	EmpIdentityRepo employee_repo.EmployeeIdentityRepositoryInterface
	EmpAuthRepo     employee_repo.EmployeeAuthRepositoryInterface
}

func bootstrapHR(db *gorm.DB) *hrBootstrap {
	empIdentityRepo := employee_repo.NewEmployeeIdentityRepository(db)
	empAuthRepo := employee_repo.NewEmployeeAuthRepository(db)
	empOnboardingRepo := employee_repo.NewEmployeeOnboardingRepository(db)
	empVerificationRepo := employee_repo.NewEmployeeVerificationRepo(db)
	empProfileRepo := employee_repo.NewEmployeeProfileRepo(db)
	empEmploymentRepo := employee_repo.NewEmployeeEmploymentRepo(db)

	policyRepo := employee_repo.NewEmployeePolicyRepo(db)
	policyRuleRepo := employee_repo.NewEmployeePolicyRuleRepository(db)
	privilegePolicyRepo := employee_repo.NewEmployeePrivilegePolicyRepository(db)

	service := hr_employee_services.NewHrEmployeeService(
		empIdentityRepo,
		empAuthRepo,
		empOnboardingRepo,
		empVerificationRepo,
		empProfileRepo,
		empEmploymentRepo,
		policyRepo,
		policyRuleRepo,
		privilegePolicyRepo,
	)

	return &hrBootstrap{
		Service:         service,
		EmpIdentityRepo: empIdentityRepo,
		EmpAuthRepo:     empAuthRepo,
	}
}
