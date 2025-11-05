package jwtpack

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidIssuer    = errors.New("invalid token issuer")
	ErrInvalidAudience  = errors.New("invalid token audience")
	ErrInvalidTokenType = errors.New("invalid token type")
	ErrMissingToken     = errors.New("authorization token is required")
)

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Verify issuer
		if claims.Issuer != JWTIssuer {
			return nil, ErrInvalidIssuer
		}

		// Verify expiration
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return nil, ErrExpiredToken
		}

		return claims, nil
	}

	return nil, ErrInvalidToken
}

func ValidateTokenType(tokenString string, expectedType string) (*Claims, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != expectedType {
		return nil, fmt.Errorf("%w: expected %s, got %s", ErrInvalidTokenType, expectedType, claims.TokenType)
	}

	return claims, nil
}

func ExtractOrgID(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if claims.UserDetail == nil || claims.UserDetail.OrgDetail == nil {
		return "", errors.New("organization detail not found in token")
	}

	return claims.UserDetail.OrgDetail.SharpnrID, nil
}
