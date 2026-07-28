package coreEmployeeServices

import (
	"context"
	"errors"
	"time"

	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Failure modes of a refresh attempt. Callers map these onto HTTP responses,
// and the distinction matters to the client: an expired or revoked session
// means "log in again", whereas reuse means the token was compromised and
// every session for that employee has just been killed.
var (
	ErrMobileSessionNotFound = errors.New("mobile session not found")
	ErrMobileSessionRevoked  = errors.New("mobile session revoked")
	ErrMobileSessionExpired  = errors.New("mobile session expired")
	ErrRefreshTokenReuse     = errors.New("refresh token reuse detected")
)

// StartMobileSession opens a new refresh-token family for a login and returns
// the identifiers to embed in the refresh token.
//
// The session expiry is taken from the JWT manager rather than a separate
// setting so the row and the token it authorises always expire together.
func (s *employeeService) StartMobileSession(
	ctx context.Context,
	employeeID uuid.UUID,
	orgID uuid.UUID,
	role string,
) (sessionID uuid.UUID, jti uuid.UUID, err error) {

	now := time.Now().UTC()
	sessionID = uuid.New()
	jti = uuid.New()

	session := &coreEmployeeModels.EmployeeRefreshSession{
		ID:             sessionID,
		EmployeeID:     employeeID,
		OrganizationID: orgID,
		CurrentJTI:     jti,
		Role:           role,
		ExpiresAt:      now.Add(s.jwtManager.RefreshTTL()),
		CreatedAt:      now,
	}

	if err := s.employeeRepo.CreateRefreshSession(ctx, session); err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return sessionID, jti, nil
}

/* ---------------------------------- <***> --------------------------------- */

// RotateMobileSession exchanges a presented refresh token for the next one in
// the family, returning the new jti and the role to stamp on the new access
// token.
//
// The rotation itself is a single conditional UPDATE, so concurrent refreshes
// carrying the same jti cannot both succeed — exactly one wins and the rest
// fall through to classification below.
//
// Losing that UPDATE is not automatically an attack. The session may simply
// have been revoked or expired between issuing and redeeming. Only a live
// session whose stored jti has moved past the presented one indicates replay:
// that token was already exchanged, so two parties hold it. In that case every
// session for the employee is revoked, which deliberately signs out the real
// user too — they must re-authenticate, and the attacker's stolen token dies
// with the family.
func (s *employeeService) RotateMobileSession(
	ctx context.Context,
	sessionID uuid.UUID,
	presentedJTI uuid.UUID,
) (newJTI uuid.UUID, role string, err error) {

	now := time.Now().UTC()
	newJTI = uuid.New()

	rotated, err := s.employeeRepo.RotateRefreshSession(ctx, sessionID, presentedJTI, newJTI, now)
	if err != nil {
		return uuid.Nil, "", err
	}

	// Read back either way: on success for the role, on failure to classify.
	session, err := s.employeeRepo.GetRefreshSessionByID(ctx, sessionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, "", ErrMobileSessionNotFound
		}
		return uuid.Nil, "", err
	}

	if rotated {
		return newJTI, session.Role, nil
	}

	switch {
	case session.RevokedAt != nil:
		return uuid.Nil, "", ErrMobileSessionRevoked

	case !session.ExpiresAt.After(now):
		return uuid.Nil, "", ErrMobileSessionExpired

	default:
		// The session is live, yet the UPDATE matched nothing — the only
		// remaining explanation is that its stored jti has moved past the one
		// presented, so this token was already exchanged. Revoke the family.
		//
		// Deliberately the default rather than an explicit jti comparison: any
		// state where a live session refuses a rotation is, by elimination,
		// a token that is no longer current, and failing closed is the right
		// default for the one branch that indicates compromise.
		if _, revokeErr := s.employeeRepo.RevokeAllRefreshSessionsForEmployee(
			ctx,
			session.EmployeeID,
			coreEmployeeModels.RevokeReasonTokenReuse,
			now,
		); revokeErr != nil {
			return uuid.Nil, "", revokeErr
		}
		return uuid.Nil, "", ErrRefreshTokenReuse
	}
}

/* ---------------------------------- <***> --------------------------------- */

// EndMobileSession revokes a single session — the logout path.
func (s *employeeService) EndMobileSession(
	ctx context.Context,
	sessionID uuid.UUID,
) error {
	return s.employeeRepo.RevokeRefreshSession(
		ctx,
		sessionID,
		coreEmployeeModels.RevokeReasonLogout,
		time.Now().UTC(),
	)
}

/* ---------------------------------- <***> --------------------------------- */

// EndAllMobileSessions revokes every live session for an employee —
// "sign out everywhere", and the same operation used after a password change.
func (s *employeeService) EndAllMobileSessions(
	ctx context.Context,
	employeeID uuid.UUID,
) (int64, error) {
	return s.employeeRepo.RevokeAllRefreshSessionsForEmployee(
		ctx,
		employeeID,
		coreEmployeeModels.RevokeReasonLogoutAll,
		time.Now().UTC(),
	)
}

/* ---------------------------------- <***> --------------------------------- */

// PurgeExpiredMobileSessions deletes session rows that have passed their
// expiry. These rows do not self-expire the way the Redis-backed keys do, so
// without a sweep the table grows once per login forever.
func (s *employeeService) PurgeExpiredMobileSessions(
	ctx context.Context,
) (int64, error) {
	return s.employeeRepo.PurgeExpiredRefreshSessions(ctx, time.Now().UTC())
}
