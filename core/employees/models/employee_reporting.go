package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 3️⃣ REPORTING HIERARCHY
// -------------------------------
type EmployeeReporting struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`
	ManagerID  uuid.UUID `gorm:"type:uuid;index"`

	Level int `gorm:"type:smallint"` // 1=Manager, 2=Skip-level

	CreatedAt time.Time
}

func (EmployeeReporting) TableName() string {
	return "employee.employee_reporting"
}

/* ---------------------------------- <***> --------------------------------- */
