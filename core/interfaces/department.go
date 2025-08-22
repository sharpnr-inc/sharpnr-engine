package interfaces

// This is Interface For Operation in Department Table
// Its directly talk and edits The Table
type DepartmentRepository interface {
	// Create Deaprtment in the Table in MainDB(Postgresql)
	// Example uses
	CreateDepartment()
	//-----------------------------------------------------------------------------------------//
	//This Function is used for Adding Removing From the DB
	AddEmployeeInDepartment()

	FindDepartmentByOrgId()

	FindEmployeeByOrgId()
}

type DepartmentService interface {
	AddDepartmentInOrg()
	RemoveDepartmentInOrg()
}
