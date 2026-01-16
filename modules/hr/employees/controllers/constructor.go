package employee_controllers

import hr_employee_services "sharpnr-engine/modules/hr/employees/services"

type employeeController struct {
	hrEmpSvc hr_employee_services.EmployeeServiceInterface
}

func NewHREmployeeController(hrEmpSvc hr_employee_services.EmployeeServiceInterface) *employeeController {
	return &employeeController{
		hrEmpSvc: hrEmpSvc,
	}
}
