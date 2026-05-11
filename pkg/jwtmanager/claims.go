// claims.go

package jwt_manager

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ---------- Access Token ----------

// Access Claims is the Structs for the token for the access token
// This is the Stucture
type AccessClaims struct {
	UserID string `json:"uid"`
	OrgID  string `json:"oid"` // SharpnrID (public)
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

// ---------- Refresh Token ----------

type RefreshClaims struct {
	UserID    string    `json:"uid"`
	SessionID string    `json:"sid"`
	OrgID     uuid.UUID `json:"org_id"`

	jwt.RegisteredClaims
}

//---------Web cookie Token -----------

type WebCookieTokenClaims struct {
	UserID string `json:"uid"`
	OrgID  string `json:"orgId"`
	jwt.RegisteredClaims
}

// ---------- Signup Token ----------

type SignupClaims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}
