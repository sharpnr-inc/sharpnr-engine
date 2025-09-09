package employeerepo

import (
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) *employeeRepository {
	return &employeeRepository{
		db: database.DB,
	}

}
