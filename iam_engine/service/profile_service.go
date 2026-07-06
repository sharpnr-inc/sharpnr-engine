package iamEngineService

import (
	"context"

	iamEngineDtos "sharpnr-engine/iam_engine/dtos"

	"github.com/google/uuid"
)

func (s *iamEngineServiceConstructor) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*iamEngineDtos.ProfileResponse_dto, error) {

	empProfile, err := s.employeeService.GetEmployeeProfile(ctx, orgID, empID)
	if err != nil {
		return nil, err
	}

	return &iamEngineDtos.ProfileResponse_dto{
		FirstName: empProfile.FirstName,
		LastName:  empProfile.LastName,
		Email:     empProfile.Email,
		Address:   empProfile.Address,
		DOB:       empProfile.DOB,
		Gender:    empProfile.Gender,
	}, nil
}
