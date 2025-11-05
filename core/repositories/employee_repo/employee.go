package employeerepo

import (
	"context"
	"fmt"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *employeeRepository) CreateEmployee(ctx context.Context, emp *models.Employee) (*models.Employee, error) {
	// Convert org string to uuid.UUID

	// Insert into database
	if err := r.db.WithContext(ctx).Create(emp).Error; err != nil {
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	return emp, nil
}

// FindEmployeeById retrieves an employee by organization ID and employee ID
func (r *employeeRepository) FindEmployeeById(ctx context.Context, orgID, empID uuid.UUID) (*models.Employee, bool, error) {
	var emp models.Employee
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
func (r *employeeRepository) FindByEmail(ctx context.Context, email string) (*models.Employee, bool, error) {
	var emp models.Employee
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
func (r *employeeRepository) ListEmployees(ctx context.Context, orgID uuid.UUID) ([]*models.Employee, error) {
	var employees []*models.Employee
	if err := r.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&employees).Error; err != nil {
		return nil, fmt.Errorf("failed to list employees: %w", err)
	}

	return employees, nil
}

// UpdateEmployeeDetails updates the provided employee's details
func (r *employeeRepository) UpdateEmployeeDetails(ctx context.Context, emp *models.Employee) (*models.Employee, error) {
	if emp.ID == uuid.Nil {
		return nil, fmt.Errorf("employee ID is required")
	}

	if err := r.db.WithContext(ctx).Save(emp).Error; err != nil {
		return nil, fmt.Errorf("failed to update employee: %w", err)
	}

	return emp, nil
}

// RemoveEmployee deletes an employee from the given organization
func (r *employeeRepository) RemoveEmployee(ctx context.Context, orgID, empID uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Where("id = ? AND organization_id = ?", empID, orgID).
		Delete(&models.Employee{}).Error; err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	return nil
}
