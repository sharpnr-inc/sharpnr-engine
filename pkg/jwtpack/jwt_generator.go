package jwtpack

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Load secret once at startup
var jwtSecret = loadJWTSecret()

// loadJWTSecret loads the JWT secret from environment (required for security)
func loadJWTSecret() []byte {
	_ = godotenv.Load() // ignore error if already loaded elsewhere
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Panic("❌ JWT_SECRET not set! Please define it in your environment or .env file.")
	}
	return []byte(secret)
}

// GenerateJwtToken creates a JWT for a given email and token type
func GenerateJwtToken(email string, tokenType string, orgID string, orgType string) (string, error) {
	claims := Claims{
		UserDetail: &UserDetail{
			OrgDetail: &OrgDetail{
				SharpnrID: orgID,
				OrgType:   orgType,
			},
			Email: email, // can be empty for OTP stage
			ID:    "",    // empty if user not created yet
			Role:  "",
		},
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // default 15m expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    JWTIssuer,
			Subject:   email,
			Audience:  jwt.ClaimStrings{JWTAudience},
		},
	}

	// Customize expiry based on token type
	switch tokenType {
	case TokenTypeCompleteRegistration:
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(30 * time.Minute))
	case TokenTypeAuth:
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	case TokenTypeVerifyOTP:
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign with secret
	return token.SignedString(jwtSecret)
}
