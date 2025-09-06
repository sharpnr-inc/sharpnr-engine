// utils/jwt.go
package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getJWTSecret() []byte {
	_ = godotenv.Load() // ignore error, assume env already loaded in main
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-super-secure-secret" // fallback, only for dev
	}
	return []byte(secret)
}

var jwtSecret = getJWTSecret()

func GenerateOTPToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(), // 15 min expiry
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func GenerateAuthToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email":     email,
		"tokenType": "auth",
		"exp":       time.Now().Add(24 * time.Hour).Unix(), // 24h expiry
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})
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
