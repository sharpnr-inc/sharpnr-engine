package employee_services

import (
	"context"
	"errors"

	employee_contracts "sharpnr-engine/core/employees/contracts"
	"sharpnr-engine/shared/utils"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

func (r *employeeService) AuthenticateEmployee(
	ctx context.Context,
	email string,
	password string,
) (*employee_contracts.EmployeeReader, error) {

	// 1. Fetch employee by email
	empl, err := r.employeeRepo.GetEmployeeByEmail(ctx, email)
	if err != nil {
		// Do NOT expose whether email exists
		return nil, ErrInvalidCredentials
	}

	// 2. Fetch auth credentials
	authEmployee, err := r.employeeRepo.GetAuthByEmployeeID(ctx, empl.ID)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// 3. Defensive check
	if authEmployee.PasswordHash == "" {
		return nil, ErrInvalidCredentials
	}

	// 4. Verify password
	if !utils.CheckPasswordHash(password, authEmployee.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// 5. Map to reader / DTO
	reader := &employee_contracts.EmployeeReader{
		EmployeeID:     empl.ID,
		Name:           empl.FirstName,
		OrganizationID: empl.OrganizationID,
		// add other safe fields
	}

	return reader, nil
}
