package app

import (
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"

	"gorm.io/gorm"
)

type hrBootstrap struct {
	Service hr_employee_services.HrEmployeeServiceInterface

	EmpIdentityRepo hr_employee_repo.EmployeeIdentityRepositoryInterface
	EmpAuthRepo     hr_employee_repo.EmployeeAuthRepositoryInterface
}

func bootstrapHR(db *gorm.DB) *hrBootstrap {
	empIdentityRepo := hr_employee_repo.NewEmployeeIdentityRepository(db)
	empAuthRepo := hr_employee_repo.NewEmployeeAuthRepository(db)
	empOnboardingRepo := hr_employee_repo.NewEmployeeOnboardingRepository(db)
	empVerificationRepo := hr_employee_repo.NewEmployeeVerificationRepo(db)
	empProfileRepo := hr_employee_repo.NewEmployeeProfileRepo(db)
	empEmploymentRepo := hr_employee_repo.NewEmployeeEmploymentRepo(db)

	policyRepo := hr_employee_repo.NewEmployeePolicyRepo(db)
	policyRuleRepo := hr_employee_repo.NewEmployeePolicyRuleRepository(db)
	privilegePolicyRepo := hr_employee_repo.NewEmployeePrivilegePolicyRepository(db)

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
