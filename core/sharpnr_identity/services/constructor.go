package sharpnr_identity_auth_services

import (
	"context"

	"github.com/google/uuid"
)

type sharpnrIdentityAuthService struct{}

// GetProfile implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) error {
	panic("unimplemented")
}

// GetRenderer implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetRenderer(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

// VerifyLoggedInEmployee implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) VerifyLoggedInEmployee(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

func NewSharpnrIdentityAuthService() *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{}
}
