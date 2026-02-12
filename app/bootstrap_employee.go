package app

import (
	employee_repo "sharpnr-engine/modules/hr/employees/repo"
	employee_services "sharpnr-engine/modules/hr/employees/services"

	"gorm.io/gorm"
)

func bootstrapEmployee(db *gorm.DB) employee_services.EmployeeServiceInterface {
	empIdentityRepo := employee_repo.NewEmployeeIdentityRepository(db)
	empAuthRepo := employee_repo.NewEmployeeAuthRepository(db)
	empOnboardingRepo := employee_repo.NewEmployeeOnboardingRepository(db)
	empVerificationRepo := employee_repo.NewEmployeeVerificationRepo(db)
	empProfileRepo := employee_repo.NewEmployeeProfileRepo(db)
	empEmploymentRepo := employee_repo.NewEmployeeEmploymentRepo(db)

	return employee_services.NewEmployeeService(
		empIdentityRepo,
		empAuthRepo,
		empOnboardingRepo,
		empVerificationRepo,
		empProfileRepo,
		empEmploymentRepo,
	)

}
