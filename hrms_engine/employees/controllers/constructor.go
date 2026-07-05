package employee_controllers

import (
	hrmsEmployeeServices "sharpnr-engine/hrms_engine/employees/services"
)

type hrmsEmployeeController struct {
	hrmsEmpService hrmsEmployeeServices.HRMSEmployeeServiceInterface
}

func NewEmployeeController(hrmsEmpService hrmsEmployeeServices.HRMSEmployeeServiceInterface) *hrmsEmployeeController {
	return &hrmsEmployeeController{
		hrmsEmpService: hrmsEmpService,
	}
}
