// manager.go

package jwt_manager

import (
	"time"

	"github.com/google/uuid"
)

type Manager interface {
	// Issue
	IssueAccess(userID, orgID, role string) (string, error)
	IssueRefresh(userID string, orgID uuid.UUID, sessionID string) (string, error)
	IssueSignup(userID string) (string, error)
	IssueWebAuthCookieToken(userID, orgID uuid.UUID) (string, error)

	// Verify
	VerifyAccess(token string) (*AccessClaims, error)
	VerifyRefresh(token string) (*RefreshClaims, error)
	VerifySignup(token string) (*SignupClaims, error)
	VerifyWebCookie(token string) (*WebCookieTokenClaims, error)
}

type manager struct {
	cfg Config
}

func New(cfg Config) Manager {
	return &manager{cfg: cfg}
}

func (m *manager) now() time.Time {
	return time.Now()
}
