package database

import (
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"gorm.io/gorm"
)

func migrateEmployee(db *gorm.DB) error {
	return db.AutoMigrate(
		&employee_models.Employee{},
		&employee_models.EmployeeOnboarding{},
		&employee_models.EmployeeVerification{},
		&employee_models.EmployeeAuth{},
		&employee_models.EmployeeSession{},
		&employee_models.EmployeeProfile{},
		&employee_models.EmployeeSystemAccess{},
		&employee_models.EmployeeBankAccount{},
		&employee_models.EmployeeDocument{},
		&employee_models.EmployeeReporting{},
		&employee_models.EmployeeEmployment{},
		&employee_models.EmployeeCompensation{},
		&employee_models.EmployeeSalaryComponent{},
		&employee_models.EmployeeLeaveType{},
		&employee_models.EmployeeLeaveBalance{},
		&employee_models.EmployeeLeaveRequest{},
	)
}
