package signupservice

import (
	"sharpnr-engine/core/interfaces"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	policyengine "sharpnr-engine/policy_engine"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"
)

// Here i want to implement or use the Organization REpository interface to interact with

// SignupService handles organization signup logic
type signupService struct {
	orgRepo interfaces.OrganizationRepository // interface, not concrete struct
	// This is Employee Repo for the creation of employee during the signup of Organization
	empRepo interfaces.EmployeeRepository

	policyRepo policyengine.PolicyRepository

	uiRenderRepo sharpnruirenderer.UIRendererRepository

	hotelRepo hotelrepo.HotelRepository
}

func NewSignupService(repo interfaces.OrganizationRepository, emprepo interfaces.EmployeeRepository, policyrepo policyengine.PolicyRepository, uiRenderRepo sharpnruirenderer.UIRendererRepository, hotelRepo hotelrepo.HotelRepository) *signupService {
	return &signupService{
		orgRepo:      repo,
		empRepo:      emprepo,
		policyRepo:   policyrepo,
		uiRenderRepo: uiRenderRepo,
		hotelRepo:    hotelRepo,
	}
}
