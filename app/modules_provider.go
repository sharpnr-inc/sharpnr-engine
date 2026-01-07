package app

import hr_employee_services "sharpnr-engine/modules/hr/employees/services"

type ModulesProvider interface {
	HrService() hr_employee_services.HrEmployeeServiceInterface
}
