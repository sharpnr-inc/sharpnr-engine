package jwt_manager

import (
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
)

func testManager(t *testing.T) Manager {
	t.Helper()

	mgr, err := New(Config{
		Secret:            "test-secret-that-is-at-least-32-bytes-long",
		Issuer:            "sharpnr-engine-test",
		Audience:          "sharpnr-client-test",
		AccessTTL:         15 * time.Minute,
		RefreshTTL:        168 * time.Hour,
		SignupTTL:         24 * time.Hour,
		WebCookieTokenTTL: 10 * time.Hour,
	})
	require.NoError(t, err)

	return mgr
}

// The session and token identifiers must survive a full sign/verify round
// trip — rotation authorises entirely on these two values, so a marshalling
// bug here would silently break reuse detection rather than fail loudly.
func TestIssueRefresh_CarriesSessionAndJTI(t *testing.T) {
	mgr := testManager(t)

	userID, orgID := uuid.New(), uuid.New()
	sessionID, jti := uuid.New(), uuid.New()

	token, err := mgr.IssueRefresh(userID, orgID, sessionID, jti)
	require.NoError(t, err)

	claims, err := mgr.VerifyRefresh(token)
	require.NoError(t, err)

	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, orgID, claims.OrgID)
	assert.Equal(t, sessionID, claims.SessionID)

	gotJTI, err := claims.JTI()
	require.NoError(t, err)
	assert.Equal(t, jti, gotJTI)
}

// Each issuance must be distinguishable, otherwise a rotated token would still
// match the stored jti and replay would go undetected.
func TestIssueRefresh_DistinctJTIProducesDistinctToken(t *testing.T) {
	mgr := testManager(t)

	userID, orgID, sessionID := uuid.New(), uuid.New(), uuid.New()
	firstJTI, secondJTI := uuid.New(), uuid.New()

	first, err := mgr.IssueRefresh(userID, orgID, sessionID, firstJTI)
	require.NoError(t, err)

	second, err := mgr.IssueRefresh(userID, orgID, sessionID, secondJTI)
	require.NoError(t, err)

	assert.NotEqual(t, first, second, "rotated tokens must differ")

	firstClaims, err := mgr.VerifyRefresh(first)
	require.NoError(t, err)
	secondClaims, err := mgr.VerifyRefresh(second)
	require.NoError(t, err)

	firstID, err := firstClaims.JTI()
	require.NoError(t, err)
	secondID, err := secondClaims.JTI()
	require.NoError(t, err)

	assert.NotEqual(t, firstID, secondID)

	// The family is stable across rotation — that is what lets one revocation
	// kill every token descended from a single login.
	assert.Equal(t, firstClaims.SessionID, secondClaims.SessionID)
}

// A refresh token minted before rotation existed carries no jti. It must be
// rejected rather than treated as jti == uuid.Nil, which would collide with
// any session whose stored value was somehow zero.
func TestRefreshClaims_JTIRejectsMissingAndMalformed(t *testing.T) {
	_, err := RefreshClaims{}.JTI()
	assert.ErrorIs(t, err, ErrInvalidToken, "empty jti must be rejected")

	malformed := RefreshClaims{}
	malformed.ID = "not-a-uuid"
	_, err = malformed.JTI()
	assert.ErrorIs(t, err, ErrInvalidToken, "malformed jti must be rejected")
}

// A refresh token must never be accepted by another verifier. The subject
// check is what stops a long-lived refresh token being replayed as an access
// token against the mobile middleware.
func TestRefreshToken_RejectedByOtherVerifiers(t *testing.T) {
	mgr := testManager(t)

	token, err := mgr.IssueRefresh(uuid.New(), uuid.New(), uuid.New(), uuid.New())
	require.NoError(t, err)

	_, err = mgr.VerifyAccess(token)
	assert.ErrorIs(t, err, ErrInvalidToken)

	_, err = mgr.VerifySignup(token)
	assert.ErrorIs(t, err, ErrInvalidToken)

	_, err = mgr.VerifyWebCookie(token)
	assert.ErrorIs(t, err, ErrInvalidToken)
}

// RefreshTTL is what stamps the session row's expiry. If it drifted from the
// configured value, rows would outlive or predecease their tokens.
func TestRefreshTTL_MatchesConfig(t *testing.T) {
	mgr := testManager(t)
	assert.Equal(t, 168*time.Hour, mgr.RefreshTTL())
}
