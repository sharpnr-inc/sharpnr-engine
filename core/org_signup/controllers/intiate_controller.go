package org_signup_controllers

import (
	"errors"
	"log/slog"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	org_signup_services "sharpnr-engine/core/org_signup/services"
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

	// This Controller First need to check Request is Correct and comes properly
	var req org_signup_dtos.OrganizationRegistrationDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	slug, err := c.signUpSvc.InitiateRegistration(ctx.Request.Context(), req)
	if err != nil {
		// Only the "already taken" cases are the caller's fault. Anything else
		// is ours (database, Redis, hashing) and must not be reported as a 400
		// with the raw error text in the body.
		switch {
		case errors.Is(err, org_signup_services.ErrUserAlreadyExists),
			errors.Is(err, org_signup_services.ErrOrgAlreadyExists):
			utils.Failure(ctx, http.StatusConflict, utils.CodeConflict, err)
		default:
			slog.Error("signup initiation failed",
				"error", err,
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			utils.Failure(
				ctx,
				http.StatusInternalServerError,
				utils.CodeInternalError,
				errors.New("could not start signup, please try again"),
			)
		}
		return
	}

	// Without this cookie every following signup step is locked out, so a
	// failure here has to surface instead of returning a cheerful 200.
	jwt, err := c.jwtMgr.IssueSignup(slug)
	if err != nil {
		slog.Error("failed to issue signup token", "error", err, "slug", slug)
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			errors.New("could not start signup, please try again"),
		)
		return
	}

	if err := c.cookies.Set(ctx, cookieManager.Signup, jwt); err != nil {
		slog.Error("failed to set signup cookie", "error", err, "slug", slug)
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			errors.New("could not start signup, please try again"),
		)
		return
	}

	utils.Success(ctx, http.StatusOK, "Otp Sent successfully")
}
