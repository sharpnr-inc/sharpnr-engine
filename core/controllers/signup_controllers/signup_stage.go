package signupcontrollers

import (
	"sharpnr-engine/core/constants"

	"github.com/gin-gonic/gin"
)

func (c *SignUpController) SignupStatus(ctx *gin.Context) {
	// Check fully authenticated first
	if _, err := ctx.Cookie(constants.TOKEN_KEY_AUTHENTICATED); err == nil {
		ctx.JSON(200, gin.H{"status": constants.SIGNUP_STAGE_AUTHENTICATED})
		return
	}

	// Check OTP verified
	if _, err := ctx.Cookie(constants.TOKEN_KEY_VERIFY_OTP); err == nil {
		ctx.JSON(200, gin.H{"status": constants.SIGNUP_STAGE_OTP_VERIFIED})
		return
	}

	// Check initial registration done
	if _, err := ctx.Cookie(constants.TOKEN_KEY_REGISTRATION); err == nil {
		ctx.JSON(200, gin.H{"status": constants.SIGNUP_STAGE_REGISTERED})
		return
	}

	// No cookie found → user not started signup
	ctx.JSON(200, gin.H{"status": constants.SIGNUP_STAGE_NONE})
}
