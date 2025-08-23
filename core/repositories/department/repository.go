package repositories

import (
	"context"
	"sharpnr-engine/core/models"

	"gorm.io/gorm"
)

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{db: db}
}

// This is Function To Create Department
func (r *departmentRepository) CreateDepartment(ctx context.Context, Dept *models.Department) error {
	if err := r.db.Create(Dept).Error; err != nil {
		return err
	}
	return nil
}

// func (r *departmentRepository) FindByID() {

// }
