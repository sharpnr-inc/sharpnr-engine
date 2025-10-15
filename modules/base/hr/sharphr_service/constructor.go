package sharphrservice

import sharphrinterface "sharpnr-engine/modules/base/hr/sharphr_interface"

type sharpHrService struct {
	employeeRepo sharphrinterface.SharpHrEmployeeRepository
}

func NewSharpHREmployeeService(repo sharphrinterface.SharpHrEmployeeRepository) *sharpHrService {
	return &sharpHrService{
		employeeRepo: repo,
	}
}
