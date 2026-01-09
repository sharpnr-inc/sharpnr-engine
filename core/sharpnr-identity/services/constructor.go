package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr-identity/dtos"
	sharpnr_identity_auth_repositories "sharpnr-engine/core/sharpnr-identity/repositories"

	"github.com/google/uuid"
)

type sharpnrIdentityAuthService struct {
	identityRepo sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface
}

// GetProfile implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) error {
	panic("unimplemented")
}

// GetRenderer implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetRenderer(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

// LoginUser implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) LoginUser(ctx context.Context, orgID uuid.UUID, req sharpnr_identity_auth_dtos.LoginRequest_dto) (sharpnr_identity_auth_dtos.WebLoginResponse_dto, error) {
	panic("unimplemented")
}

// VerifyLoggedInEmployee implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) VerifyLoggedInEmployee(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

func NewSharpnrIdentityAuthService(
	identityRepo sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		identityRepo: identityRepo,
	}
}
