package employee_controllers

import employee_services "sharpnr-engine/core/employees/services"

type employeeController struct {
	hrEmpSvc employee_services.EmployeesHRServiceInterface
}

func NewEmployeeController(hrEmpSvc employee_services.EmployeesHRServiceInterface) *employeeController {
	return &employeeController{
		hrEmpSvc: hrEmpSvc,
	}
}
