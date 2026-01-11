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

type ModulesProvider interface {
	OrgSignUpService() org_signup_services.OrgSignUpServiceInterface
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	HospitalityHotelService() hotelservices.HotelService
	HospitalityBarService() barservices.BarServices
	HospitalityKitchenService() kitchenservices.KitchenServiceInterface
	HospitalityMenuService() menuservice.MenuService
	HospitalityRestaurantService() restaurantservices.RestaurantService

	HrService() hr_employee_services.HrEmployeeServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
}
