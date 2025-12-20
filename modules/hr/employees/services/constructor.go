package hr_employee_services

import hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"

type hrEmployeeService struct {
	hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface
}

func NewHrEmployeeService(hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface) *hrEmployeeService {
	return &hrEmployeeService{
		hrEmpRepo: hrEmpRepo,
	}
}
