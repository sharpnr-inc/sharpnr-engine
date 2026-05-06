package employee_repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	employee_models "sharpnr-engine/modules/hr/employees/models"
)

type employeeRepo struct {
	db *gorm.DB
}

// CreateEmplVerification implements [EmployeeRepositoryInterface].
func (r *employeeRepo) CreateEmplVerification(ctx context.Context, v *employee_models.EmployeeVerification) error {
	panic("unimplemented")
}

// UpdateIdentity implements [EmployeeRepositoryInterface].
func (r *employeeRepo) UpdateIdentity(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]interface{}) error {
	panic("unimplemented")
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}
