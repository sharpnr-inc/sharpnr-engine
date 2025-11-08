package sharphrservice

import (
	sharphrinterface "sharpnr-engine/modules/hr/sharphr_interface"
	policyengine "sharpnr-engine/policy_engine"
)

type sharpHrService struct {
	employeeRepo sharphrinterface.SharpHrEmployeeRepository
	policyRepo   policyengine.PolicyRepository
}

func NewSharpHREmployeeService(repo sharphrinterface.SharpHrEmployeeRepository, policyRepo policyengine.PolicyRepository) *sharpHrService {
	return &sharpHrService{
		employeeRepo: repo,
		policyRepo:   policyRepo,
	}
}
