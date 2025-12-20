package hr_employee_repo

import (
	"context"
	"fmt"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hrEmployeeRepo struct {
	db *gorm.DB
}

func NewHrEmployeeRepository(db *gorm.DB) *hrEmployeeRepo {
	return &hrEmployeeRepo{
		db: db,
	}
}

func (r *hrEmployeeRepo) CreateEmployee(ctx context.Context, emp *sharphr_models.Employee) (*sharphr_models.Employee, error) {
	// Convert org string to uuid.UUID

	// Insert into database
	if err := r.db.WithContext(ctx).Create(emp).Error; err != nil {
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	return emp, nil
}

// FindEmployeeById retrieves an employee by organization ID and employee ID
func (r *hrEmployeeRepo) FindEmployeeById(ctx context.Context, orgID, empID uuid.UUID) (*sharphr_models.Employee, bool, error) {
	var emp sharphr_models.Employee
	err := r.db.WithContext(ctx).
		Where("id = ? AND organization_id = ?", empID, orgID).
		First(&emp).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed to fetch employee: %w", err)
	}

	return &emp, true, nil
}

// FindByEmail retrieves an employee by organization ID and email
func (r *hrEmployeeRepo) FindByEmail(ctx context.Context, email string) (*sharphr_models.Employee, bool, error) {
	var emp sharphr_models.Employee
	err := r.db.WithContext(ctx).
		Where("email = ? ", email).
		First(&emp).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed to fetch employee by email: %w", err)
	}

	return &emp, true, nil
}

// ListEmployees retrieves all employees for a given organization
func (r *hrEmployeeRepo) ListEmployees(ctx context.Context, orgID uuid.UUID) ([]*sharphr_models.Employee, error) {
	var employees []*sharphr_models.Employee
	if err := r.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&employees).Error; err != nil {
		return nil, fmt.Errorf("failed to list employees: %w", err)
	}

	return employees, nil
}

// UpdateEmployeeDetails updates the provided employee's details
func (r *hrEmployeeRepo) UpdateEmployeeDetails(ctx context.Context, emp *sharphr_models.Employee) (*sharphr_models.Employee, error) {
	if emp.ID == uuid.Nil {
		return nil, fmt.Errorf("employee ID is required")
	}

	if err := r.db.WithContext(ctx).Save(emp).Error; err != nil {
		return nil, fmt.Errorf("failed to update employee: %w", err)
	}

	return emp, nil
}

// RemoveEmployee deletes an employee from the given organization
func (r *hrEmployeeRepo) RemoveEmployee(ctx context.Context, orgID, empID uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Where("id = ? AND organization_id = ?", empID, orgID).
		Delete(&sharphr_models.Employee{}).Error; err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	return nil
}
