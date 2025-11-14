package signupservice

import (
	"sharpnr-engine/core/interfaces"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	policyengine "sharpnr-engine/policy_engine"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"
)

// Here i want to implement or use the Organization REpository interface to interact with

// SignupService handles organization signup logic
type signupService struct {
	orgRepo interfaces.OrganizationRepositoryInterface // interface, not concrete struct
	// This is Employee Repo for the creation of employee during the signup of Organization

	hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface

	policyRepo policyengine.PolicyRepository

	uiRenderRepo sharpnruirenderer.UIRendererRepository

	hotelRepo hotelrepo.HotelRepository

	barRepo barrepo.BarRepository

	restaurantRepo restaurantrepo.RestaurantRepository

	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewSignupService(repo interfaces.OrganizationRepositoryInterface, hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface, policyrepo policyengine.PolicyRepository, uiRenderRepo sharpnruirenderer.UIRendererRepository, hotelRepo hotelrepo.HotelRepository, barRepo barrepo.BarRepository, restaurantRepo restaurantrepo.RestaurantRepository, kitchenRepo kitchenrepo.KitchenRepositoryInterface) *signupService {
	return &signupService{
		orgRepo:        repo,
		hrEmpRepo:      hrEmpRepo,
		policyRepo:     policyrepo,
		uiRenderRepo:   uiRenderRepo,
		hotelRepo:      hotelRepo,
		barRepo:        barRepo,
		restaurantRepo: restaurantRepo,
		kitchenRepo:    kitchenRepo,
	}
}
