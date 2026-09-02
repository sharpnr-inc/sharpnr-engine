package orgDepartmentRepo

import (
	"gorm.io/gorm"
)

type orgDepartmentRepository struct {
	db *gorm.DB
}

func NewOrgDepartmentRepository(db *gorm.DB) *orgDepartmentRepository {
	return &orgDepartmentRepository{
		db: db,
	}
}
