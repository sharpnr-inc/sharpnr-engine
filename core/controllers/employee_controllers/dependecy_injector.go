package employeecontrollers

import "sharpnr-engine/core/interfaces"

type EmployeeControllerimpl struct {
	service interfaces.EmployeeService
}

func EmployeeController(service interfaces.EmployeeService) *EmployeeControllerimpl {
	return &EmployeeControllerimpl{service: service}
}
