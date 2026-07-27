package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeDepartments struct {
	ID           uuid.UUID
	EmployeeID   uuid.UUID
	DepartmentID uuid.UUID
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
