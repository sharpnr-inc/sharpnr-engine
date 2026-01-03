package jwt_manager

import (
	"github.com/golang-jwt/jwt/v5"
)

func (m *manager) parse(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(m.cfg.Secret), nil
		},
	)

	if err != nil || !token.Valid {
		return ErrInvalidToken
	}

	// ✅ Issuer check
	issuer, err := claims.GetIssuer()
	if err != nil || issuer != m.cfg.Issuer {
		return ErrInvalidIssuer
	}

	// ✅ Audience check
	audiences, err := claims.GetAudience()
	if err != nil {
		return ErrInvalidAudience
	}

	validAudience := false
	for _, aud := range audiences {
		if aud == m.cfg.Audience {
			validAudience = true
			break
		}
	}

	if !validAudience {
		return ErrInvalidAudience
	}

	return nil
}

// ---------- Verify Access ----------

func (m *manager) VerifyAccess(token string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	return claims, nil
}

// ---------- Verify Refresh ----------

func (m *manager) VerifyRefresh(token string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	return claims, nil
}

// ---------- Verify Signup ----------

func (m *manager) VerifySignup(token string) (*SignupClaims, error) {
	claims := &SignupClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	return claims, nil
}
