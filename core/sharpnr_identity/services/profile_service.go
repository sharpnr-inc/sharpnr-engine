package sharpnr_identity_auth_services

import (
	"context"
	"errors"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"

	"github.com/google/uuid"
)

func (s *sharpnrIdentityAuthService) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*sharpnr_identity_auth_dtos.ProfileResponse_dto, error) {

	empProfile, err := s.employeeService.GetEmployeeProfile(ctx, orgID, empID)
	if err != nil {
		return nil, err
	}

	return &sharpnr_identity_auth_dtos.ProfileResponse_dto{
		Email:     empProfile.Email,
		FirstName: empProfile.FirstName,
		LastName:  empProfile.LastName,
		DOB:       empProfile.DOB,
		Address:   empProfile.Address,
		Gender:    empProfile.Gender,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sharpnrIdentityAuthService) GetAuthenticatedUserDetails(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) (sharpnr_identity_auth_dtos.UserDataResponseDto, error) {

	empl, err := s.employeeService.GetEmployeeDetails(ctx, orgID, emplID)
	if err != nil {
		return sharpnr_identity_auth_dtos.UserDataResponseDto{}, err
	}

	orgDetails, err := s.orgService.GetOrgDetailsByID(ctx, orgID, emplID)
	if err != nil {
		return sharpnr_identity_auth_dtos.UserDataResponseDto{}, err
	}

	if orgDetails == nil {
		return sharpnr_identity_auth_dtos.UserDataResponseDto{},
			errors.New("organization details not found")
	}

	out := sharpnr_identity_auth_dtos.UserDataResponseDto{
		Email:     empl.Email,
		FirstName: empl.FirstName,
		LastName:  empl.LastName,
		ID:        empl.ID,
		Role:      empl.Role,
		Org: sharpnr_identity_auth_dtos.OrgData{
			Name:    orgDetails.Name,
			SharpID: orgDetails.SharpID,
		},
	}

	return out, nil
}
