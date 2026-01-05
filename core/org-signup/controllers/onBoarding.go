package org_signup_controllers

import (
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"

	"github.com/gin-gonic/gin"
)

func (c *orgSignUpController) OnBoardOrganization(ctx *gin.Context) {
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

}
