package app

import (
	employee_sere_adapter "sharpnr-engine/modules/hr/employees/sere_adapter"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
)

func RegisterAllWorkFlow(
	sere *sere.SERE,
	hrEmployeeService hr_employee_services.HrEmployeeServiceInterface,
) {
	employee_sere_adapter.RegisterHRWorkflows(sere, hrEmployeeService)
}
