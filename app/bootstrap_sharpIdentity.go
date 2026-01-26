package app

import (
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	employee_services "sharpnr-engine/modules/hr/employees/services"

	"gorm.io/gorm"
)

func bootstrapIdentity(
	db *gorm.DB,
	employeeService employee_services.EmployeeServiceInterface,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {

	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(employeeService)
}
