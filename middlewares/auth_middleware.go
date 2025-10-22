package middlewares

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/pkg/jwtpack"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ClaimsContextKey = "jwt_claims"
	UserIDContextKey = "user_id"
	EmailContextKey  = "email"
	OrgIDContextKey  = "org_id"
	RoleContextKey   = "role"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Authorisation Header is required",
				Error: dto.ErrorResponse{
					Message: "Authorisation Failed",
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return

		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid authorization header format",
				Error: dto.ErrorResponse{
					Message: "Authorization header must be Bearer",
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}

			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		claims, err := jwtpack.ValidateToken(tokenString)
		if err != nil {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid or Expired token",
				Error: dto.ErrorResponse{
					Message: err.Error(),
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Set(ClaimsContextKey, claims)
		if claims.UserDetail != nil {
			c.Set(EmailContextKey, claims.UserDetail.Email)
			c.Set(UserIDContextKey, claims.UserDetail.ID)
			c.Set(RoleContextKey, claims.UserDetail.Role)

			if claims.UserDetail.OrgDetail != nil {
				c.Set(OrgIDContextKey, claims.UserDetail.OrgDetail.SharpnrID)
			}
		}

		c.Next()

	}
}
