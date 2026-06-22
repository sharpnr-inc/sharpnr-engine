package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/hrms_engine/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) CreateEmployee(
	ctx context.Context,
	emp *employee_models.Employee,
) error {

	if emp.ID == uuid.Nil {
		emp.ID = uuid.New()
	}

	return r.conn(ctx).WithContext(ctx).Create(emp).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetEmployeeByID(
	ctx context.Context,
	orgID uuid.UUID, empID uuid.UUID,
) (*employee_models.Employee, error) {

	var emp employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("id = ?", empID).
		First(&emp).
		Error

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetEmployeeByEmail(
	ctx context.Context,
	email string,
) (*employee_models.Employee, error) {

	var emp employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&emp).
		Error

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) EmployeeExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&employee_models.Employee{}).
		Where("email = ?", email).
		Count(&count).
		Error

	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) ListEmployeesByOrg(
	ctx context.Context,
	orgID uuid.UUID,
) ([]*employee_models.Employee, error) {

	var employees []*employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&employees).
		Error

	if err != nil {
		return nil, err
	}

	return employees, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) UpdateEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]interface{}) error {
	// 1. Ensure we are scoped to the correct Organization and Employee
	result := r.db.WithContext(ctx).
		Model(&employee_models.Employee{}).
		Where("id = ? AND organization_id = ?", emplID, orgID).
		Updates(updates)

	// 2. Handle database-level errors
	if result.Error != nil {
		return result.Error
	}

	// 3. Handle cases where the record doesn't exist or orgID mismatch
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
