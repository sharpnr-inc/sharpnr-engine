// utils/jwt.go
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your-super-secure-secret") // use env var in production

func GenerateJWTToken(orgID uuid.UUID, email string) (string, error) {
	claims := jwt.MapClaims{
		"org_id": orgID.String(),
		"email":  email,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // token expires in 24h
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWTToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
