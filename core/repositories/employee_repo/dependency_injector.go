package employeerepo

import (
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type initEmployeeRepository struct {
	db *gorm.DB
}

func EmployeeRepo(db *gorm.DB) *initEmployeeRepository {
	return &initEmployeeRepository{
		db: database.DB,
	}

}
