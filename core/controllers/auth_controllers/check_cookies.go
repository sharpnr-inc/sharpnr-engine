package authcontrollers

import (
	"net/http"
	"sharpnr-engine/core/constants"
	"sharpnr-engine/pkg/jwtpack"

	"github.com/gin-gonic/gin"
)

func (s *AuthController) CheckCookies(ctx *gin.Context) {

	token, err := ctx.Cookie(constants.TOKEN_KEY_AUTHENTICATED)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "No auth token found",
		})
		return
	}

	claims, err := jwtpack.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid or expired token",
		})
		return
	}

	// Token is valid
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Valid cookie",
		"user":    claims, // optional: return claims if needed
	})
}
