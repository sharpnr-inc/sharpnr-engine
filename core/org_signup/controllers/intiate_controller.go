package org_signup_controllers

import (
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"

	"github.com/gin-gonic/gin"
)

func (c *orgSignUpController) InitiateNewOrgRegistration(ctx *gin.Context) {

	// Step 1. Verify Request body and Check if All fields are present
	var signupRequest org_signup_dtos.InitiateOrgSignUpRequest_Dto

	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Step 2: Call service layer
	signupSessionID, err := c.orgSvc.InitiateSignup(
		ctx.Request.Context(),
		signupRequest,
	)
	if err != nil {
		// map domain errors
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	signupToken, _ := c.jwtMgr.IssueSignup(signupSessionID)

	c.cookies.SetSignup(ctx, signupToken)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "OTP sent to registered email",
	})

}
