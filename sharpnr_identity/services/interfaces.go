package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_dtos "sharpnr-engine/sharpnr_identity/dtos"

	"github.com/google/uuid"
)

type SharpnrIdentityAuthServiceInterface interface {
	LoginUser(
		ctx context.Context,
		req sharpnr_identity_auth_dtos.LoginRequest_dto) (
		*sharpnr_identity_auth_dtos.WebLoginResponse_dto,
		error)

	GetProfile(
		ctx context.Context,
		orgID uuid.UUID,
		empID uuid.UUID,
	) (*sharpnr_identity_auth_dtos.ProfileResponse_dto, error)

	GetAuthenticatedUserDetails(ctx context.Context, orgID, emplID uuid.UUID) (sharpnr_identity_auth_dtos.UserDataResponseDto, error)
	UpdateSuperAdminDetails(ctx context.Context, orgID, empID uuid.UUID, input sharpnr_identity_auth_dtos.UpdateSuperAdminDetailsInput) error
}
