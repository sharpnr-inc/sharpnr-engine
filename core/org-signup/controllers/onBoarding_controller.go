package org_signup_controllers

import (
	"fmt"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
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

	var onBoardingRequest org_signup_dtos.OrgOnBoardingRequest_dto

	if err := ctx.ShouldBindJSON(&onBoardingRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	res, err := c.orgSvc.StartOrgOnBoarding(ctx.Request.Context(), userID, onBoardingRequest)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.EmployeeID)

	// c.jwtMgr.IssueAccess()
	ctx.JSON(http.StatusOK, gin.H{"done": "done"})
	c.cookies.ClearSignup(ctx)
	// c.jwtMgr.IssueRefresh()
	// c.jwtMgr.IssueAccess()
	// c.cookies.Set(ctx, cookieManager.Auth)

}
