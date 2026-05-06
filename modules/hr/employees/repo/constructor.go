package employee_repo

import (
	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}
