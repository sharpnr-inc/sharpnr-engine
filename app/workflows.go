package app

import (
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	"sharpnr-engine/modules/accounting/acc_sere_wf_adapters"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_service"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	bar_sere_adapters "sharpnr-engine/modules/hospitality/bar/sere_adapters"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotel_sere_adapter "sharpnr-engine/modules/hospitality/hotel/sere_adapter"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchen_sere_adapters "sharpnr-engine/modules/hospitality/kitchen/sere_adapters"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	menu_sere_adapter "sharpnr-engine/modules/hospitality/menu/sere_adapter"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
	restaurant_sere_adapters "sharpnr-engine/modules/hospitality/restaurant/sere_adapters"
	employee_sere_adapter "sharpnr-engine/modules/hr/employees/sere_adapter"
	employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
)

func RegisterAllWorkFlow(
	sere *sere.SERE,

	hrEmployeeService employee_services.EmployeeServiceInterface,
	sharpnrIdSvc sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	hotelRepo hotelrepo.HotelRepositoryInterface,
	hotelService hotelservices.HotelService,
	barRepo barrepo.BarRepository,
	barService barservices.BarServices,
	restaurantRepo restaurantrepo.RestaurantRepository,
	restuarantService restaurantservices.RestaurantService,

	menuRepo menurepo.MenuRepositoryInterface,
	menuService menuservice.MenuService,
	accountingRepo accounting_repo.AccountingRepositoryInterface,
	accountingService accounting_service.AccountingServiceInterface,

) {
	employee_sere_adapter.RegisterHREmployeeWorkflows(sere, hrEmployeeService)

	kitchen_sere_adapters.RegisterKitchenWorkFlow(sere, kitchenRepo)

	hotel_sere_adapter.RegisterHotelWorkFlow(sere, hotelRepo)
	bar_sere_adapters.RegisterBarWorkFlow(sere, barRepo)
	restaurant_sere_adapters.RegisterRestaurantWorkFlow(sere, restaurantRepo)

	menu_sere_adapter.RegisterHospitalityBarMenuWorkflow(sere, menuRepo)
	menu_sere_adapter.RegisterHospitalityHotelMenuWorkflow(sere, menuRepo)
	menu_sere_adapter.RegisterHospitalityRestaurantMenuWorkflow(sere, menuRepo)

	acc_sere_wf_adapters.RegisterAccountingBarWorkflows(sere, accountingRepo)
	acc_sere_wf_adapters.RegisterAccountingHotelWorkflow(sere, accountingRepo)
	acc_sere_wf_adapters.RegisterAccountingRestaurantWorkflow(sere, accountingRepo)
}
