package sharpnr_identity_auth_services

import (
	"context"
	"errors"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"
	"sharpnr-engine/shared/utils"
)

// LoginUser implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) LoginUser(
	ctx context.Context,
	req sharpnr_identity_auth_dtos.LoginRequest_dto,
) (*sharpnr_identity_auth_dtos.WebLoginResponse_dto, error) {

	authenticatedEmp, err := s.empIdentity.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if authenticatedEmp.Email != req.Email {
		return nil, errors.New("Wrong Email or No User Found")
	}

	empAuthDetails, err := s.employeeAuth.GetByEmployeeID(ctx, authenticatedEmp.ID)
	if err != nil {
		return nil, err
	}

	passMatched := utils.CheckPasswordHash(req.Password, empAuthDetails.PasswordHash)
	if !passMatched {
		return nil, errors.New("Wrong Password")
	}

	res := sharpnr_identity_auth_dtos.WebLoginResponse_dto{
		OrgId:      authenticatedEmp.OrganizationID,
		Email:      authenticatedEmp.Email,
		EmployeeID: authenticatedEmp.ID,
		Name:       authenticatedEmp.FirstName,
	}

	return &res, nil
}
