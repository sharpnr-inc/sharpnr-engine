package org_signup_controllers

import (
	"errors"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	organizationErrs "sharpnr-engine/core/organization/errs"
	"sharpnr-engine/pkg/cookieManager"
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
// @Failure      409  {object}  utils.Response "Email already registered"
// @Failure      500  {object}  utils.Response "Internal Error"
// @Router       /signup/initiate [post]
func (c *orgSignUpController) InitiateNewOrgRegistration(ctx *gin.Context) {

	req, ok := utils.BindJSON[org_signup_dtos.OrganizationRegistrationDTO](ctx)
	if !ok {
		return
	}

	slug, err := c.signUpSvc.InitiateRegistration(ctx.Request.Context(), *req)
	if err != nil {
		switch {
		case errors.Is(err, organizationErrs.OrganizationAlreadyExistsErr),
			errors.Is(err, organizationErrs.OrganizationAlreadyExistsErr):
			utils.Failure(ctx, http.StatusConflict, err)
		default:
			utils.Failure(
				ctx,
				http.StatusInternalServerError,
				errors.New("could not start signup, please try again"))
		}
		return
	}

	// Without this cookie every following signup step is locked out, so a
	// failure here has to surface instead of returning a cheerful 200.
	jwt, err := c.jwtMgr.IssueSignup(slug)
	if err != nil {
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			errors.New("could not start signup, please try again"))

		return
	}

	if err := c.cookies.Set(ctx, cookieManager.Signup, jwt); err != nil {
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			errors.New("could not start signup, please try again"))

		return
	}

	utils.Success(ctx, http.StatusOK, "Otp Sent successfully")
}
