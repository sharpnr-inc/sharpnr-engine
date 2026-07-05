package sharpnr_identity_auth_services

import (
	"context"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"

	"github.com/google/uuid"
)

func (s *sharpnrIdentityAuthService) UpdateSuperAdminDetails(ctx context.Context, orgID, empID uuid.UUID, input sharpnr_identity_auth_dtos.UpdateSuperAdminDetailsInput) error {

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
