package org_signup_controllers

import (
	"errors"
	"log/slog"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	organization_services "sharpnr-engine/core/organization/services"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/pkg/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {

	// 1. Read signup context using the strongly-typed helper
	slugId, exist := webMiddleware.GetSignupUserID(ctx)
	if !exist || slugId == uuid.Nil {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
		return
	}

	// 2. Bind request
	var req org_signup_dtos.VerifyOtpRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// 3. Verify OTP & create org + employee
	orgID, empID, err := c.signUpSvc.VerifyOtp(
		ctx.Request.Context(),
		req.Otp,
		slugId,
	)

	if err != nil {
		// A bad OTP is a 401; a failure to build the organization is not. Both
		// used to come back as "unauthorized", which told the user their code
		// was wrong when the real problem was ours.
		switch {
		case errors.Is(err, organization_services.ErrInvalidOTP),
			errors.Is(err, organization_services.ErrSignupNotFound),
			errors.Is(err, organization_services.ErrInvalidSignupData):
			slog.Info("signup otp rejected", "reason", err.Error(), "slug", slugId)
			utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, err)
		default:
			slog.Error("signup verification failed",
				"error", err,
				"slug", slugId,
				"ip", ctx.ClientIP(),
			)
			utils.Failure(
				ctx,
				http.StatusInternalServerError,
				utils.CodeInternalError,
				errors.New("could not complete signup, please try again"),
			)
		}
		return
	}

	// 4. Issue JWT
	token, err := c.jwtMgr.IssueWebAuthCookieToken(
		empID,
		orgID,
	)

	if err != nil {
		slog.Error("failed to issue auth token after signup",
			"error", err,
			"org_id", orgID,
			"employee_id", empID,
		)
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			errors.New("signup completed but session could not be started, please sign in"),
		)
		return
	}

	// 5. Cleanup signup cookie & set auth cookie
	c.cookies.Clear(ctx, cookieManager.Signup)
	c.cookies.Set(ctx, cookieManager.Auth, token)

	// The CSRF cookie is deliberately not HttpOnly so the frontend can echo it
	// back in X-CSRF-Token. Putting the auth JWT in it would hand that session
	// token to any script on the page and undo the HttpOnly flag on the auth
	// cookie. Mint an unrelated random value, exactly as WebAuthMiddleware does.
	c.cookies.Set(ctx, cookieManager.CSRF, uuid.New().String())

	slog.Info("signup completed", "org_id", orgID, "employee_id", empID)

	// 6. Success response
	utils.Success(
		ctx,
		http.StatusOK,
		"Signup verified successfully",
		gin.H{
			"organization_id": orgID,
			"employee_id":     empID,
		},
	)
}
