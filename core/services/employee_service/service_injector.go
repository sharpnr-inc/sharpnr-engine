package employeeservice

import "sharpnr-engine/core/interfaces"

type EmployeeServiceImpl struct {
	emplRepo interfaces.EmployeeRepository
	orgRepo  interfaces.OrganizationRepository
}

// EmployeeService is a constructor that injects the repositories
func EmployeeService(orgRepo interfaces.OrganizationRepository, emplRepo interfaces.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		orgRepo:  orgRepo,
		emplRepo: emplRepo,
	}
}
