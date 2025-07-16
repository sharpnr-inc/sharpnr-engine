package interfaces

import (
	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

type DepartmentService interface {
	CreateDefaultAccountingDepartment(orgUUID uuid.UUID) (*models.Department, error)
	// Add other department-related methods as needed
}

type DepartmentRepository interface {
	//This is Function Initialize the Base Departments
	CreateBaseDeparment()

	// this is To check if Department ID exists
	FindDepartmentById()

	//
	IfDepartmentIDExist()
}
