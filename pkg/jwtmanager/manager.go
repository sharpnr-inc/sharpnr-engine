package jwt_manager

import "time"

type Manager interface {
	// Issue
	IssueAccess(userID, orgID, role string) (string, error)
	IssueRefresh(userID, sessionID string) (string, error)
	IssueSignup(userID string) (string, error)

	// Verify
	VerifyAccess(token string) (*AccessClaims, error)
	VerifyRefresh(token string) (*RefreshClaims, error)
	VerifySignup(token string) (*SignupClaims, error)
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
