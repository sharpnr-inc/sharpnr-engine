package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr-identity/dtos"

	"github.com/google/uuid"
)

type SharpnrIdentityAuthServiceInterface interface {
	// This Service method Authorizes the user
	// This is the common
	LoginUser(
		ctx context.Context,
		req sharpnr_identity_auth_dtos.LoginRequest_dto) (
		*sharpnr_identity_auth_dtos.WebLoginResponse_dto,
		error)

	GetRenderer(
		ctx context.Context,
		orgID uuid.UUID,
		employeeID uuid.UUID,
	) error

	GetProfile(
		ctx context.Context,
		orgID uuid.UUID,
		empID uuid.UUID,
	) error

	VerifyLoggedInEmployee(
		ctx context.Context,
		orgID uuid.UUID,
		employeeID uuid.UUID,
	) error
}
