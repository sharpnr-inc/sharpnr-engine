package authenticationEngineServices

import (
	"context"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"

	"github.com/google/uuid"
)

func (s *authenticationService) UpdateSuperAdminDetails(ctx context.Context, orgID, empID uuid.UUID, input authenticationEngineDtos.UpdateSuperAdminDetailsInput) error {

	// Check if the secondary email already exists for another user
	_, err := s.employeeService.IsUserExists(ctx, input.SecondaryEmail)
	if err != nil {
		return err
	}

	if err := s.employeeService.UpdateSuperAdminEmployee(ctx, orgID, empID, &coreEmployeeContracts.UpdateSuperAdminEmployeeInput{
		OrganizationID: input.OrgID,
		EmployeeID:     &input.EmployeeID,
		FirstName:      &input.FirstName,
		LastName:       &input.LastName,
		SecondaryEmail: &input.SecondaryEmail,
	}); err != nil {
		return err
	}

	return nil
}
