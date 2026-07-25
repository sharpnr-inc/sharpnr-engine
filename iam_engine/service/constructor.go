package iamEngineService

import coreEmployeeServices "sharpnr-engine/core/employees/services"

type iamEngineServiceConstructor struct {
	employeeService coreEmployeeServices.EmployeeIamEngineServiceInterface
}

func NewIamEngineServiceConstructor(employeeService coreEmployeeServices.EmployeeIamEngineServiceInterface) *iamEngineServiceConstructor {
	return &iamEngineServiceConstructor{
		employeeService: employeeService,
	}
}
