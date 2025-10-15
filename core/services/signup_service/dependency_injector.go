package signupservice

import (
	"sharpnr-engine/core/interfaces"
	policyengine "sharpnr-engine/policy_engine"
)

// Here i want to implement or use the Organization REpository interface to interact with

// SignupService handles organization signup logic
type signupService struct {
	orgRepo interfaces.OrganizationRepository // interface, not concrete struct
	// This is Employee Repo for the creation of employee during the signup of Organisation
	empRepo interfaces.EmployeeRepository

	policyRepo policyengine.PolicyRepository
}

func NewSignupService(repo interfaces.OrganizationRepository, emprepo interfaces.EmployeeRepository, policyrepo policyengine.PolicyRepository) *signupService {
	return &signupService{
		orgRepo:    repo,
		empRepo:    emprepo,
		policyRepo: policyrepo,
	}
}
