package app

import (
	bar_sere_adapters "sharpnr-engine/modules/hospitality/bar/sere_adapters"
	hotel_sere_adapter "sharpnr-engine/modules/hospitality/hotel/sere_adapter"
	kitchen_sere_adapters "sharpnr-engine/modules/hospitality/kitchen/sere_adapters"
	menu_sere_adapter "sharpnr-engine/modules/hospitality/menu/sere_adapter"
	restaurant_sere_adapters "sharpnr-engine/modules/hospitality/restaurant/sere_adapters"
	employee_sere_adapter "sharpnr-engine/modules/hr/employees/sere_adapter"
	"sharpnr-engine/sere"
	acc_sere_wf_adapters "sharpnr-engine/service_engines/accounting_engine/workflow_adapter"
)

func RegisterAllWorkFlow(
	sere *sere.SERE,

) {
	employee_sere_adapter.RegisterHREmployeeWorkflows(sere)

	kitchen_sere_adapters.RegisterKitchenWorkFlow(sere)

	hotel_sere_adapter.RegisterHotelWorkFlow(sere)
	bar_sere_adapters.RegisterBarWorkFlow(sere)
	restaurant_sere_adapters.RegisterRestaurantWorkFlow(sere)

	menu_sere_adapter.RegisterHospitalityBarMenuWorkflow(sere)
	menu_sere_adapter.RegisterHospitalityHotelMenuWorkflow(sere)
	menu_sere_adapter.RegisterHospitalityRestaurantMenuWorkflow(sere)

	acc_sere_wf_adapters.RegisterAccountingBarWorkflows(sere)
	acc_sere_wf_adapters.RegisterAccountingHotelWorkflow(sere)
	acc_sere_wf_adapters.RegisterAccountingRestaurantWorkflow(sere)
}
