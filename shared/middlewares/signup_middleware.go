package middlewares

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/pkg/jwtpack"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireTokenType(tokenType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Authorization header is required",
				Error: dto.ErrorResponse{
					Message: "Authorization header is required",
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}
			ctx.JSON(http.StatusUnauthorized, res)
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid authorization header format",
				Error: dto.ErrorResponse{
					Message: "Authorization header must be 'Bearer {token}'",
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}

			ctx.JSON(http.StatusUnauthorized, res)
			ctx.Abort()
			return
		}
		tokenString := tokenParts[1]

		claims, err := jwtpack.ValidateTokenType(tokenString, tokenType)
		if err != nil {
			res := dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "",
				Error: dto.ErrorResponse{
					Message: "",
					IsError: true,
				},
				Data:      nil,
				Timestamp: time.Now(),
				Token:     "",
			}

			ctx.JSON(http.StatusUnauthorized, res)
			ctx.Abort()
			return
		}

		ctx.Set(ClaimsContextKey, claims)
		if claims.UserDetail != nil {
			ctx.Set(EmailContextKey, claims.UserDetail.Email)
			ctx.Set(UserIDContextKey, claims.UserDetail.ID)
			ctx.Set(RoleContextKey, claims.UserDetail.Role)

			if claims.UserDetail.OrgDetail != nil {
				ctx.Set(OrgIDContextKey, claims.UserDetail.OrgDetail.SharpnrID)
			}
		}

		ctx.Next()

	}
}
