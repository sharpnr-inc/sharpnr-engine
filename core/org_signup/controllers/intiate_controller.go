package org_signup_controllers

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// InitiateNewOrgRegistration godoc
// @Summary      Initiate Organization Registration
// @Description  Validates the request and starts the registration process.
// @Tags         signup
// @Accept       json
// @Produce      json
// @Param        registration  body      org_signup_dtos.OrganizationRegistrationDTO  true  "Details"
// @Success      200  {object}  utils.Response{data=string} "Successfully initiated registration"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Router       /signup/initiate [post]
func (c *orgSignUpController) InitiateNewOrgRegistration(ctx *gin.Context) {

	// This Controller First need to check Request is Correct and comes properly
	var req org_signup_dtos.OrganizationRegistrationDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	slug, err := c.signUpSvc.InitiateRegistration(ctx.Request.Context(), req)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Signup error", err)
		return
	}

	jwt, _ := c.jwtMgr.IssueSignup(slug)
	c.cookies.Set(ctx, cookieManager.Signup, jwt)

	utils.Success(ctx, 200, "Otp Sent successfully")

}
