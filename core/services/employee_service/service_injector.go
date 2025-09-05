package employeeservice

import "sharpnr-engine/core/interfaces"

type EmployeeServiceImpl struct {
	emplRepo interfaces.EmployeeRepository
	orgRepo  interfaces.IOrgRegistrationRepository
}

// EmployeeService is a constructor that injects the repositories
func EmployeeService(orgRepo interfaces.IOrgRegistrationRepository, emplRepo interfaces.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		orgRepo:  orgRepo,
		emplRepo: emplRepo,
	}
}
