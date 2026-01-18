package org_signup_controllers

import (
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// ! Swagger is Remaining
// response it returns is nothing only err or 200 response
// This part of the OrgSignUp Module
func (c *orgSignUpController) InitiateNewOrgRegistration(ctx *gin.Context) {

	// This Controller First need to check Request is Correct and comes properly
	var req org_signup_dtos.OrganizationRegistrationDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	c.signUpSvc.InitiateRegistration(ctx.Request.Context(), req)

}
