package signupservice

import "sharpnr-engine/core/interfaces"

// Here i want to implement or use the Organization REpository interface to interact with

// SignupService handles organization signup logic
type signupService struct {
	orgRepo interfaces.OrganizationRepository // interface, not concrete struct
}

func NewSignupService(repo interfaces.OrganizationRepository) *signupService {
	return &signupService{
		orgRepo: repo,
	}
}
