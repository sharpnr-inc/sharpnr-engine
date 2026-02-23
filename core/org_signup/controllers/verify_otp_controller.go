package org_signup_controllers

import (
	"errors"
	"net/http"
	"sharpnr-engine/cookieManager"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {

	// 1. Read signup context
	slugId, exist := ctx.Get(webMiddleware.SessionIDContextKey)
	if !exist {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
		return
	}

	// 2. Type assert
	userIDStr, ok := slugId.(string)
	if !ok || userIDStr == "" {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("invalid signup context"),
		)
		return
	}

	// 3. Parse UUID
	slug, err := uuid.Parse(userIDStr)
	if err != nil || slug == uuid.Nil {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("invalid signup context"),
		)
		return
	}

	// 4. Bind request
	var req org_signup_dtos.VerifyOtpRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// 5. Verify OTP & create org + employee
	orgID, empID, err := c.signUpSvc.VerifyOtp(ctx, req.Otp, slug)
	if err != nil {
		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			err,
		)
		return
	}

	// 6. Issue JWT
	token, err := c.jwtMgr.IssueWebAuthCookieToken(empID, orgID)
	if err != nil {
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			err,
		)
		return
	}

	// 7. Cleanup signup cookie & set auth cookie
	c.cookies.ClearSignup(ctx)
	c.cookies.Set(ctx, cookieManager.Auth, token)

	// 8. Success response (FIXED)
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

func (c *orgSignUpController) VerifySlug(ctx *gin.Context) {

	userIDVal, exists := ctx.Get(webMiddleware.SignupUserIDContextKey)
	if !exists {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
		return
	}

	// Type assert to uuid.UUID
	userIDStr, ok := userIDVal.(string)
	if !ok || userIDStr == "" {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("invalid signup context"),
		)
		return
	}

	slug, err := uuid.Parse(userIDStr)
	if err != nil || slug == uuid.Nil {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("invalid signup context"),
		)
		return
	}

	// Call service
	email, err := c.signUpSvc.VerifyInitiationRegCookie(
		ctx.Request.Context(),
		slug,
	)
	if err != nil {
		utils.Failure(
			ctx,
			http.StatusBadRequest,
			utils.CodeUnauthorized,
			err,
		)
		return
	}

	utils.Success(ctx, http.StatusOK, "Signup slug verified", email)
}
