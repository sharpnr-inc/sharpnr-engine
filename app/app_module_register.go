package app

import (
	org_signup_services "sharpnr-engine/core/org-signup/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
)

type AppModules struct {
	orgSignupService org_signup_services.OrgSignUpServiceInterface
	sharpAuthService sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	hospitality_hotelService      hotelservices.HotelService
	hospitality_barService        barservices.BarServices
	hospitality_restaurantService restaurantservices.RestaurantService
	hospitality_kitchenService    kitchenservices.KitchenServiceInterface
	hospitality_menuService       menuservice.MenuService

	hrService         hr_employee_services.HrEmployeeServiceInterface
	accountingService accounting_service.AccountingServiceInterface
}

// AccountingService implements ModulesProvider.
func (m *AppModules) AccountingService() accounting_service.AccountingServiceInterface {
	return m.accountingService
}

// HospitalityBarService implements ModulesProvider.
func (m *AppModules) HospitalityBarService() barservices.BarServices {
	return m.hospitality_barService
}

// HospitalityHotelService implements ModulesProvider.
func (m *AppModules) HospitalityHotelService() hotelservices.HotelService {
	return m.hospitality_hotelService
}

// HospitalityKitchenService implements ModulesProvider.
func (m *AppModules) HospitalityKitchenService() kitchenservices.KitchenServiceInterface {
	return m.hospitality_kitchenService
}

// HospitalityMenuService implements ModulesProvider.
func (m *AppModules) HospitalityMenuService() menuservice.MenuService {
	return m.hospitality_menuService
}

// HospitalityRestaurantService implements ModulesProvider.
func (m *AppModules) HospitalityRestaurantService() restaurantservices.RestaurantService {
	return m.hospitality_restaurantService
}

// SharpAuthService implements ModulesProvider.
func (m *AppModules) SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return m.sharpAuthService
}

func (m *AppModules) HrService() hr_employee_services.HrEmployeeServiceInterface {
	return m.hrService
}

func (m *AppModules) OrgSignUpService() org_signup_services.OrgSignUpServiceInterface {
	return m.orgSignupService
}
