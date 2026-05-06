package app

import (
	employee_repo "sharpnr-engine/modules/hr/employees/repo"
	employee_services "sharpnr-engine/modules/hr/employees/services"

	"gorm.io/gorm"
)

func bootstrapEmployeeIdentity(db *gorm.DB) employee_services.EmployeesIdentityServiceInterface {
	employeeRepo := employee_repo.NewEmployeeRepository(db)
	return employee_services.NewEmployeeService(
		employeeRepo,
	)
}

func bootstrapEmployeeHR(db *gorm.DB) employee_services.EmployeesHRServiceInterface {
	employeeRepo := employee_repo.NewEmployeeRepository(db)
	return employee_services.NewEmployeeService(employeeRepo)
}
