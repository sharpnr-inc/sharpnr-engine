package interfaces

import (
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"github.com/google/uuid"
)

type DepartmentService interface {
	CreateDefaultAccountingDepartment(orgUUID uuid.UUID) (*models.Department, error)
	// Add other department-related methods as needed
}