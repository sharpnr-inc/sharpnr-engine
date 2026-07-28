package authenticationEngineServices

import (
	"context"
	"errors"
	"fmt"
	"time"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// Re-exported so controllers can classify a refresh failure without importing
// the employee services package directly.
var (
	ErrMobileSessionNotFound = coreEmployeeServices.ErrMobileSessionNotFound
	ErrMobileSessionRevoked  = coreEmployeeServices.ErrMobileSessionRevoked
	ErrMobileSessionExpired  = coreEmployeeServices.ErrMobileSessionExpired
	ErrRefreshTokenReuse     = coreEmployeeServices.ErrRefreshTokenReuse
)

// StartMobileSession opens a refresh-token family for a fresh mobile login.
func (s *authenticationService) StartMobileSession(
	ctx context.Context,
	employeeID uuid.UUID,
	orgID uuid.UUID,
	role string,
) (sessionID uuid.UUID, jti uuid.UUID, err error) {
	return s.employeeService.StartMobileSession(ctx, employeeID, orgID, role)
}

/* ---------------------------------- <***> --------------------------------- */

// RotateMobileSession exchanges a refresh token for the next one in its family.
//
// employeeID and orgID come from the verified token claims and are used only
// to attribute the audit entry written when replay is detected — the rotation
// itself is authorised entirely by sessionID and presentedJTI against the
// stored session.
func (s *authenticationService) RotateMobileSession(
	ctx context.Context,
	sessionID uuid.UUID,
	presentedJTI uuid.UUID,
	employeeID uuid.UUID,
	orgID uuid.UUID,
) (newJTI uuid.UUID, role string, err error) {

	newJTI, role, err = s.employeeService.RotateMobileSession(ctx, sessionID, presentedJTI)

	// Replay means a refresh token was presented twice. Every session for the
	// employee has just been revoked, which is a security event worth a
	// permanent record — it is the only signal that a device token leaked.
	if errors.Is(err, ErrRefreshTokenReuse) {
		s.recordRefreshReuse(ctx, employeeID, orgID, sessionID)
	}

	return newJTI, role, err
}

// recordRefreshReuse writes the audit entry for a detected replay. Audit
// failures are swallowed deliberately: the caller is already rejecting the
// request and revoking the family, and losing the log entry must not turn a
// successful revocation into a 500 that leaves the client retrying.
func (s *authenticationService) recordRefreshReuse(
	ctx context.Context,
	employeeID uuid.UUID,
	orgID uuid.UUID,
	sessionID uuid.UUID,
) {
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	_ = s.auditEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &employeeID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         "authentication",
		Action:         "employee.refresh_token_reuse",
		ResourceType:   "employee",
		ResourceID:     &employeeID,
		Description: fmt.Sprintf(
			"Refresh token replay detected on %s; all mobile sessions revoked",
			time.Now().UTC().Format("2006-01-02 15:04:05 MST"),
		),
		Metadata: map[string]any{
			"session_id": sessionID.String(),
		},
	})
}

/* ---------------------------------- <***> --------------------------------- */

// EndMobileSession revokes one session — the logout path.
func (s *authenticationService) EndMobileSession(
	ctx context.Context,
	sessionID uuid.UUID,
) error {
	return s.employeeService.EndMobileSession(ctx, sessionID)
}

/* ---------------------------------- <***> --------------------------------- */

// PurgeExpiredMobileSessions is driven by the background sweeper in main.go.
func (s *authenticationService) PurgeExpiredMobileSessions(
	ctx context.Context,
) (int64, error) {
	return s.employeeService.PurgeExpiredMobileSessions(ctx)
}
