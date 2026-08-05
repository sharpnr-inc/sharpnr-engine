package authenticationEngineServices

import (
	"context"
	"errors"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"

	"github.com/google/uuid"
)

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
			errors.New("organization details not found")
	}

	out := authenticationEngineDtos.UserDataResponseDto{
		Email:     empl.Email,
		FirstName: empl.FirstName,
		LastName:  empl.LastName,
		ID:        empl.ID,
		Org: authenticationEngineDtos.OrgData{
			Name:    orgDetails.Name,
			SharpID: orgDetails.SharpID,
		},
	}

	return out, nil
}
