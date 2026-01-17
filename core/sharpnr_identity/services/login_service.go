package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"
)

// LoginUser implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) LoginUser(
	ctx context.Context,
	req sharpnr_identity_auth_dtos.LoginRequest_dto,
) (*sharpnr_identity_auth_dtos.WebLoginResponse_dto, error) {
	return nil, nil
}
