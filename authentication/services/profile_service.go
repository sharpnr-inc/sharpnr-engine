package authenticationEngineServices

import (
	"context"
	"errors"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"

	"github.com/google/uuid"
)

func (s *authenticationService) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*authenticationEngineDtos.ProfileResponse_dto, error) {

	empProfile, err := s.employeeService.GetEmployeeProfile(ctx, orgID, empID)
	if err != nil {
		return nil, err
	}

	return &authenticationEngineDtos.ProfileResponse_dto{
		Email:     empProfile.Email,
		FirstName: empProfile.FirstName,
		LastName:  empProfile.LastName,
		DOB:       empProfile.DOB,
		Address:   empProfile.Address,
		Gender:    empProfile.Gender,
	}, nil
}

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
