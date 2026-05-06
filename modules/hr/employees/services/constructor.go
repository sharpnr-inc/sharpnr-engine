package employee_services

import (
	employee_repo "sharpnr-engine/modules/hr/employees/repo"
)

type employeeService struct {
	employeeRepo employee_repo.EmployeeRepositoryInterface
}

func NewEmployeeService(
	employeeRepo employee_repo.EmployeeRepositoryInterface,
) *employeeService {

	return &employeeService{
		employeeRepo: employeeRepo,
	}
}
