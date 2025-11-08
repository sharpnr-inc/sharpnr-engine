package sharphrinterface

import sharphrdtos "sharpnr-engine/modules/hr/sharphr_dtos"

type SharpHrEmployeeRepository interface {
	FindEmployeeByEmail()

	CreateEmployeeInOrg()

	GetEmployeeDetailsByEmployeeID()
}

type SharpHrEmployeeService interface {
	AddEmployeeInOrg() error

	GetEmployeeDetailById() error

	ListAllEmployeeInOrg() error

	GetEmployeeDetailByEmail(sharpId, query_email string, org_uuid any) (*sharphrdtos.CheckIfEmployeeExist_RESPONSE_DTO, error)
}
