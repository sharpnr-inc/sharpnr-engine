package coreEmployeeServices

import (
	"context"
	"testing"
	"time"

	coreEmployeeModels "sharpnr-engine/core/employees/models"
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// fakeSessionRepo implements only the refresh-session methods. The embedded
// interface is nil, so any other repo call panics loudly rather than silently
// returning a zero value.
type fakeSessionRepo struct {
	coreEmployeeRepo.EmployeeRepositoryInterface

	session *coreEmployeeModels.EmployeeRefreshSession

	revokedAllFor    []uuid.UUID
	revokedAllReason string
	created          *coreEmployeeModels.EmployeeRefreshSession
}

func (f *fakeSessionRepo) CreateRefreshSession(
	_ context.Context,
	session *coreEmployeeModels.EmployeeRefreshSession,
) error {
	f.created = session
	f.session = session
	return nil
}

func (f *fakeSessionRepo) GetRefreshSessionByID(
	_ context.Context,
	sessionID uuid.UUID,
) (*coreEmployeeModels.EmployeeRefreshSession, error) {
	if f.session == nil || f.session.ID != sessionID {
		return nil, gorm.ErrRecordNotFound
	}
	return f.session, nil
}

// RotateRefreshSession mirrors the WHERE clause of the real conditional
// UPDATE. Keeping the guards identical is the point: the service relies on
// this returning false for a superseded jti so it can classify the failure.
func (f *fakeSessionRepo) RotateRefreshSession(
	_ context.Context,
	sessionID, expectedJTI, newJTI uuid.UUID,
	now time.Time,
) (bool, error) {
	s := f.session
	if s == nil || s.ID != sessionID {
		return false, nil
	}
	if s.CurrentJTI != expectedJTI || s.RevokedAt != nil || !s.ExpiresAt.After(now) {
		return false, nil
	}

	s.CurrentJTI = newJTI
	s.RotatedAt = &now
	return true, nil
}

func (f *fakeSessionRepo) RevokeAllRefreshSessionsForEmployee(
	_ context.Context,
	employeeID uuid.UUID,
	reason string,
	at time.Time,
) (int64, error) {
	f.revokedAllFor = append(f.revokedAllFor, employeeID)
	f.revokedAllReason = reason

	if f.session != nil && f.session.EmployeeID == employeeID && f.session.RevokedAt == nil {
		f.session.RevokedAt = &at
		f.session.RevokedReason = reason
		return 1, nil
	}
	return 0, nil
}

func newServiceWithSession(
	t *testing.T,
	session *coreEmployeeModels.EmployeeRefreshSession,
) (*employeeService, *fakeSessionRepo) {
	t.Helper()

	mgr, err := jwt_manager.New(jwt_manager.Config{
		Secret:            "test-secret-that-is-at-least-32-bytes-long",
		Issuer:            "sharpnr-engine-test",
		Audience:          "sharpnr-client-test",
		AccessTTL:         15 * time.Minute,
		RefreshTTL:        168 * time.Hour,
		SignupTTL:         24 * time.Hour,
		WebCookieTokenTTL: 10 * time.Hour,
	})
	require.NoError(t, err)

	repo := &fakeSessionRepo{session: session}
	return &employeeService{employeeRepo: repo, jwtManager: mgr}, repo
}

func liveSession(jti uuid.UUID) *coreEmployeeModels.EmployeeRefreshSession {
	return &coreEmployeeModels.EmployeeRefreshSession{
		ID:             uuid.New(),
		EmployeeID:     uuid.New(),
		OrganizationID: uuid.New(),
		CurrentJTI:     jti,
		Role:           "ORG_ADMIN",
		ExpiresAt:      time.Now().UTC().Add(24 * time.Hour),
	}
}

/* ---------------------------------- <***> --------------------------------- */

func TestRotateMobileSession_HappyPathAdvancesJTI(t *testing.T) {
	currentJTI := uuid.New()
	session := liveSession(currentJTI)
	svc, repo := newServiceWithSession(t, session)

	newJTI, role, err := svc.RotateMobileSession(context.Background(), session.ID, currentJTI)

	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, newJTI)
	assert.NotEqual(t, currentJTI, newJTI, "rotation must issue a different jti")

	// Role survives the refresh — it used to be dropped, leaving the new
	// access token with an empty role claim.
	assert.Equal(t, "ORG_ADMIN", role)

	assert.Equal(t, newJTI, repo.session.CurrentJTI, "stored jti must advance")
	assert.Empty(t, repo.revokedAllFor, "a normal refresh must not revoke anything")
}

// The security-critical case: presenting a jti that has already been exchanged
// means two parties hold the same token. Everything for that employee dies.
func TestRotateMobileSession_ReuseRevokesWholeFamily(t *testing.T) {
	firstJTI := uuid.New()
	session := liveSession(firstJTI)
	svc, repo := newServiceWithSession(t, session)

	// Legitimate refresh moves the session on to a new jti.
	_, _, err := svc.RotateMobileSession(context.Background(), session.ID, firstJTI)
	require.NoError(t, err)

	// The attacker replays the token that was just retired.
	_, _, err = svc.RotateMobileSession(context.Background(), session.ID, firstJTI)

	assert.ErrorIs(t, err, ErrRefreshTokenReuse)
	assert.Equal(t, []uuid.UUID{session.EmployeeID}, repo.revokedAllFor,
		"reuse must revoke every session for the employee")
	assert.Equal(t, coreEmployeeModels.RevokeReasonTokenReuse, repo.revokedAllReason)
	assert.NotNil(t, repo.session.RevokedAt)
}

// After a reuse-triggered revocation, even the legitimate holder of the
// current token is locked out — that is the intended blast radius.
func TestRotateMobileSession_CurrentTokenDeadAfterReuse(t *testing.T) {
	firstJTI := uuid.New()
	session := liveSession(firstJTI)
	svc, repo := newServiceWithSession(t, session)

	secondJTI, _, err := svc.RotateMobileSession(context.Background(), session.ID, firstJTI)
	require.NoError(t, err)

	_, _, err = svc.RotateMobileSession(context.Background(), session.ID, firstJTI)
	require.ErrorIs(t, err, ErrRefreshTokenReuse)

	// The real client now tries the token it legitimately holds.
	_, _, err = svc.RotateMobileSession(context.Background(), session.ID, secondJTI)
	assert.ErrorIs(t, err, ErrMobileSessionRevoked)
	assert.NotNil(t, repo.session.RevokedAt)
}

func TestRotateMobileSession_RevokedSessionIsNotTreatedAsReuse(t *testing.T) {
	currentJTI := uuid.New()
	session := liveSession(currentJTI)
	revokedAt := time.Now().UTC().Add(-time.Minute)
	session.RevokedAt = &revokedAt
	session.RevokedReason = coreEmployeeModels.RevokeReasonLogout

	svc, repo := newServiceWithSession(t, session)

	_, _, err := svc.RotateMobileSession(context.Background(), session.ID, currentJTI)

	assert.ErrorIs(t, err, ErrMobileSessionRevoked)
	assert.Empty(t, repo.revokedAllFor,
		"an ordinary logout must not escalate to a reuse revocation")
}

func TestRotateMobileSession_ExpiredSession(t *testing.T) {
	currentJTI := uuid.New()
	session := liveSession(currentJTI)
	session.ExpiresAt = time.Now().UTC().Add(-time.Hour)

	svc, repo := newServiceWithSession(t, session)

	_, _, err := svc.RotateMobileSession(context.Background(), session.ID, currentJTI)

	assert.ErrorIs(t, err, ErrMobileSessionExpired)
	assert.Empty(t, repo.revokedAllFor)
}

func TestRotateMobileSession_UnknownSession(t *testing.T) {
	svc, _ := newServiceWithSession(t, nil)

	_, _, err := svc.RotateMobileSession(context.Background(), uuid.New(), uuid.New())

	assert.ErrorIs(t, err, ErrMobileSessionNotFound)
}

/* ---------------------------------- <***> --------------------------------- */

// The session row must expire with the token it authorises, not on some
// separate schedule.
func TestStartMobileSession_ExpiryTracksRefreshTTL(t *testing.T) {
	svc, repo := newServiceWithSession(t, nil)

	before := time.Now().UTC()
	sessionID, jti, err := svc.StartMobileSession(
		context.Background(), uuid.New(), uuid.New(), "ORG_ADMIN",
	)
	require.NoError(t, err)

	require.NotNil(t, repo.created)
	assert.Equal(t, sessionID, repo.created.ID)
	assert.Equal(t, jti, repo.created.CurrentJTI,
		"the issued jti must be the one stored, or the first refresh looks like replay")
	assert.Equal(t, "ORG_ADMIN", repo.created.Role)

	expectedExpiry := before.Add(168 * time.Hour)
	assert.WithinDuration(t, expectedExpiry, repo.created.ExpiresAt, 5*time.Second)
}
