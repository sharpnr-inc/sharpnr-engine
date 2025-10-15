package middlewares

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/pkg/jwtpack"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireTokenType(tokenType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cookieName string

		// Map token type to cookie name using constants
		switch tokenType {
		case jwtpack.TokenTypeVerifyOTP:
			cookieName = constants.TOKEN_KEY_REGISTRATION
		case jwtpack.TokenTypeCompleteRegistration:
			cookieName = constants.TOKEN_KEY_VERIFY_OTP
		case jwtpack.TokenTypeAuth:
			cookieName = constants.TOKEN_KEY_AUTHENTICATED
		default:
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token type",
				Error:      dto.ErrorResponse{Message: "Invalid token type", IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		// Read cookie
		tokenString, err := ctx.Cookie(cookieName)
		if err != nil || tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized: missing token",
				Error:      dto.ErrorResponse{Message: "Missing token in cookie", IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		// Validate JWT
		claims, err := jwtpack.ValidateTokenType(tokenString, tokenType)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
				Error:      dto.ErrorResponse{Message: err.Error(), IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		// Set claims in context
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
