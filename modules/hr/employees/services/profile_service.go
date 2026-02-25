package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *employeeService) GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*employee_contracts.EmployeeProfileOutput, error) {

	emp, err := s.empIdentityRepo.GetByID(ctx, empID)
	if err != nil {
		return nil, err
	}

	empProfile, err := s.empProfileRepo.GetByEmployeeID(ctx, empID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 3. Map Identity (Guaranteed to have data)
	output := &employee_contracts.EmployeeProfileOutput{
		FirstName: emp.FirstName,
		LastName:  emp.LastName,
		Email:     emp.Email,
	}

	// 4. Map Profile only if it was found
	if empProfile != nil {
		output.Address = empProfile.Address
		output.Gender = empProfile.Gender

		// Ensure we don't format a zero-value time
		if !empProfile.DateOfBirth.IsZero() {
			output.DOB = empProfile.DateOfBirth.Local().String()
		}
	}

	return output, nil

}
