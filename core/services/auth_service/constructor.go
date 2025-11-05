package authservice

import "sharpnr-engine/core/interfaces"

type loginService struct {
	orgRepo interfaces.OrganizationRepository
	empRepo interfaces.EmployeeRepository
}

func NewAuthService(orgrepo interfaces.OrganizationRepository, emprepo interfaces.EmployeeRepository) *loginService {
	return &loginService{
		orgRepo: orgrepo,
		empRepo: emprepo,
	}
}
