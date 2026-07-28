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

	// Mobile refresh-token sessions. Rotation and reuse-detection rules live in
	// core/employees/services/mobile_session_service.go; this layer adds the
	// audit trail for detected replays.
	StartMobileSession(ctx context.Context, employeeID, orgID uuid.UUID, role string) (sessionID, jti uuid.UUID, err error)
	RotateMobileSession(ctx context.Context, sessionID, presentedJTI, employeeID, orgID uuid.UUID) (newJTI uuid.UUID, role string, err error)
	EndMobileSession(ctx context.Context, sessionID uuid.UUID) error
	PurgeExpiredMobileSessions(ctx context.Context) (int64, error)
}
