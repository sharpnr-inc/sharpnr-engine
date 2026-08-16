package authenticationEngineServices

import (
	"context"
	"errors"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	authenticationEngineErrs "sharpnr-engine/authentication/errs"
	sentinelServiceErrs "sharpnr-engine/sentinel/errs"

	"github.com/google/uuid"
)

// Placeholder shown for an authenticated employee who has no role assigned in
// the organization yet.
const roleNotAssigned = "ROLE NOT ASSIGNED"

/* ---------------------------------- <***> --------------------------------- */

func (s *authenticationService) GetAuthenticatedUserDetails(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) (authenticationEngineDtos.UserDataResponseDto, error) {

	empl, err := s.employeeService.GetEmployeeDetails(ctx, orgID, emplID)
	if err != nil {
		return authenticationEngineDtos.UserDataResponseDto{}, err
	}

	orgDetails, err := s.orgService.GetOrgDetailsByID(ctx, orgID, emplID)
	if err != nil {
		return authenticationEngineDtos.UserDataResponseDto{}, err
	}

	if orgDetails == nil {
		return authenticationEngineDtos.UserDataResponseDto{},
			authenticationEngineErrs.OrgDetailsNotFoundErr
	}

	// An employee without a role is still a valid authenticated user, so the
	// profile is returned with a placeholder role instead of failing the request.
	roleName := roleNotAssigned

	role, err := s.sentinelService.GetEmployeeRole(ctx, orgID, emplID)
	switch {
	case err == nil:
		roleName = role.Name
	case errors.Is(err, sentinelServiceErrs.EmployeeRoleNotFoundErr):
		// keep the placeholder
	default:
		return authenticationEngineDtos.UserDataResponseDto{}, err
	}

	out := authenticationEngineDtos.UserDataResponseDto{
		Email:     empl.Email,
		FirstName: empl.FirstName,
		LastName:  empl.LastName,
		ID:        empl.ID,
		Role:      roleName,
		Org: authenticationEngineDtos.OrgData{
			Name:    orgDetails.Name,
			SharpID: orgDetails.SharpID,
		},
	}

	return out, nil
}
