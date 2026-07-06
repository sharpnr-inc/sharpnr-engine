package authenticationEngineServices

import (
	"context"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"

	"github.com/google/uuid"
)

type AuthenticationServiceInterface interface {
	LoginUser(
		ctx context.Context,
		req authenticationEngineDtos.LoginRequest_dto) (
		*authenticationEngineDtos.WebLoginResponse_dto,
		error)

	GetProfile(
		ctx context.Context,
		orgID uuid.UUID,
		empID uuid.UUID,
	) (*authenticationEngineDtos.ProfileResponse_dto, error)

	GetAuthenticatedUserDetails(ctx context.Context, orgID, emplID uuid.UUID) (authenticationEngineDtos.UserDataResponseDto, error)
}
