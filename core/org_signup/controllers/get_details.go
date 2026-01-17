package org_signup_controllers

import (
	"github.com/gin-gonic/gin"
)

func (c *orgSignUpController) GetSignUpDetails(ctx *gin.Context) {

	// // 1️⃣ Extract signup user ID from context
	// userIDVal, exists := ctx.Get(webMiddleware.SignupUserIDContextKey)
	// if !exists {
	// 	ctx.JSON(http.StatusForbidden, gin.H{
	// 		"message": "Signup context missing or expired",
	// 	})
	// 	return
	// }

	// userID, ok := userIDVal.(string)
	// if !ok || userID == "" {
	// 	ctx.JSON(http.StatusForbidden, gin.H{
	// 		"message": "Invalid signup context",
	// 	})
	// 	return
	// }

	// stagingOrg, err := c.orgSvc.VerifySlug(ctx.Request.Context(), userID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"session": "Signup Session expired"})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"Message": "Success",
	// 	"Org":     stagingOrg,
	// })
}
