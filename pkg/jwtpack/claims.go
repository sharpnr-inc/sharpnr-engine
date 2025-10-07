package jwtpack

import "github.com/golang-jwt/jwt/v5"

const (
	TokenTypeVerifyOTP            = "VERIFY_OTP"
	TokenTypeCompleteRegistration = "COMPLETE_REGISTRATION"
	TokenTypeAuth                 = "AUTH"
)

const (
	JWTIssuer   = "Sharpnr-Engine-v1"
	JWTAudience = "sharpnr-app"
)

// OrgDetail represents organization information
type OrgDetail struct {
	SharpnrID string `json:"SHARPNR_ID"`
	OrgType   string `json:"ORG_TYPE"`
}

type UserDetail struct {
	OrgDetail *OrgDetail `json:"ORG_DETAIL"`
	Email     string     `json:"Email"`
	ID        string     `json:"ID"`
	Role      string     `json:"Role"`
}

// Claims represents the JWT claims structure
type Claims struct {
	UserDetail *UserDetail `json:"USER_DETAIL,omitempty"` // pointer + omitempty
	TokenType  string      `json:"TOKEN_TYPE"`
	jwt.RegisteredClaims
}
