package jwt_manager

import (
	"github.com/golang-jwt/jwt/v5"
)

// ---------- Issue Access ----------

func (m *manager) IssueAccess(userID, orgID, role string) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		OrgID:  orgID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.cfg.Issuer,
			Audience:  []string{m.cfg.Audience},
			IssuedAt:  jwt.NewNumericDate(m.now()),
			ExpiresAt: jwt.NewNumericDate(m.now().Add(m.cfg.AccessTTL)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(m.cfg.Secret))
}

// ---------- Issue Refresh ----------

func (m *manager) IssueRefresh(userID, orgId string, sessionID string) (string, error) {
	claims := RefreshClaims{
		UserID:    userID,
		SessionID: sessionID,
		OrgID:     orgId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.cfg.Issuer,
			Audience:  []string{m.cfg.Audience},
			IssuedAt:  jwt.NewNumericDate(m.now()),
			ExpiresAt: jwt.NewNumericDate(m.now().Add(m.cfg.RefreshTTL)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(m.cfg.Secret))
}

// ---------- Issue Signup ----------

func (m *manager) IssueSignup(userID string) (string, error) {
	claims := SignupClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.cfg.Issuer,
			Audience:  []string{m.cfg.Audience},
			IssuedAt:  jwt.NewNumericDate(m.now()),
			ExpiresAt: jwt.NewNumericDate(m.now().Add(m.cfg.SignupTTL)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(m.cfg.Secret))
}
