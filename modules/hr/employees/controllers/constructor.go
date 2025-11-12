package hr_employee_controllers

import hr_employee_services "sharpnr-engine/modules/hr/employees/services"

type hrEmployeeController struct {
	hrEmpSvc hr_employee_services.HrEmployeeServiceInterface
}

func NewHREmployeeController(hrEmpSvc hr_employee_services.HrEmployeeServiceInterface) *hrEmployeeController {
	return &hrEmployeeController{
		hrEmpSvc: hrEmpSvc,
	}
}
