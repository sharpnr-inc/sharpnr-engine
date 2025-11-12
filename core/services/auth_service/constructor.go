package authservice

import (
	"sharpnr-engine/core/interfaces"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
)

type loginService struct {
	orgRepo   interfaces.OrganizationRepository
	empRepo   interfaces.EmployeeRepository
	hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface
}

func NewAuthService(orgrepo interfaces.OrganizationRepository, emprepo interfaces.EmployeeRepository, hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface) *loginService {
	return &loginService{
		orgRepo:   orgrepo,
		empRepo:   emprepo,
		hrEmpRepo: hrEmpRepo,
	}
}
