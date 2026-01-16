package org_signup_controllers

import (
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"

	"github.com/gin-gonic/gin"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {

	// 1️⃣ Extract signup user ID from context
	userIDVal, exists := ctx.Get(webMiddleware.SignupUserIDContextKey)
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Signup context missing or expired",
		})
		return
	}

	userID, ok := userIDVal.(string)
	if !ok || userID == "" {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid signup context",
		})
		return
	}

	var req org_signup_dtos.VerifyOtpRequest_Dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}

	if err := c.orgSvc.VerifyOtp(ctx.Request.Context(), userID, req); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Otp": "Doest not Match"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP verified successfully",
	})

}
