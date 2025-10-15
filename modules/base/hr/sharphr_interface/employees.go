package sharphrinterface

type SharpHrEmployeeRepository interface {
	FindEmployeeByEmail()

	CreateEmployeeInOrg()

	GetEmployeeDetailsByEmployeeID()
}

type SharpHrEmployeeService interface {
}
