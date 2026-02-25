package sharpnr_identity_auth_services

import (
	"context"
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
