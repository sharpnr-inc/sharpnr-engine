package authcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *AuthController) Logout(ctx *gin.Context) {

	// Overwrite the auth cookie with an expired one
	ctx.SetCookie(
		"_sharpnr_auth_token_", // cookie name (must match login)
		"",                     // empty value
		-1,                     // maxAge < 0 deletes the cookie
		"/",                    // path (must match login)
		"",                     // domain (same as login)
		false,                  // secure (true in prod HTTPS)
		true,                   // httpOnly
	)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})

}
