package jwt_manager

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ---------- Access Token ----------

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

// ---------- Signup Token ----------

type SignupClaims struct {
	UserID string `json:"uid"`

	jwt.RegisteredClaims
}
