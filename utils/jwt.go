package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecretKey = []byte("your-very-secret-key") // Ideally load from config/env

func GenerateToken(orgUUID uuid.UUID, orgName string) (string, error) {
	claims := jwt.MapClaims{
		"org_uuid": orgUUID.String(),
		"name":     orgName,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiry
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
