package app

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	iamEngineService "sharpnr-engine/iam_engine/service"
)

func bootstrapIamEngine(
	employeeService coreEmployeeServices.EmployeeIamEngineServiceInterface,
) iamEngineService.IamEngineServiceInterface {

	return iamEngineService.NewIamEngineServiceConstructor(employeeService)
}
