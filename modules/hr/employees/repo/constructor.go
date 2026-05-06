package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

// CreateEmployee implements [EmployeeRepositoryInterface].
func (r *employeeRepo) CreateEmployee(ctx context.Context, emp *employee_models.Employee) error {
	panic("unimplemented")
}

// EmployeeExistsByEmail implements [EmployeeRepositoryInterface].
func (r *employeeRepo) EmployeeExistsByEmail(ctx context.Context, email string) (bool, error) {
	panic("unimplemented")
}

// GetEmployeeByEmail implements [EmployeeRepositoryInterface].
func (r *employeeRepo) GetEmployeeByEmail(ctx context.Context, email string) (*employee_models.Employee, error) {
	panic("unimplemented")
}

// GetEmployeeByID implements [EmployeeRepositoryInterface].
func (r *employeeRepo) GetEmployeeByID(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*employee_models.Employee, error) {
	panic("unimplemented")
}

// ListEmployeesByOrg implements [EmployeeRepositoryInterface].
func (r *employeeRepo) ListEmployeesByOrg(ctx context.Context, orgID uuid.UUID) ([]*employee_models.Employee, error) {
	panic("unimplemented")
}

// UpdateEmployee implements [EmployeeRepositoryInterface].
func (r *employeeRepo) UpdateEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]interface{}) error {
	panic("unimplemented")
}

// CreateEmplVerification implements [EmployeeRepositoryInterface].
func (r *employeeRepo) CreateEmplVerification(ctx context.Context, v *employee_models.EmployeeVerification) error {
	panic("unimplemented")
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}
