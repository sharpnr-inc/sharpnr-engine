package org_signup_controllers

import (
	"errors"
	"log"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/pkg/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {
	log.Println("[VerifyOtpForSignup] request started")

	// 1. Read signup context using the strongly-typed helper
	slugId, exist := webMiddleware.GetSignupUserID(ctx)
	log.Printf("[VerifyOtpForSignup] signup context found=%v value=%v\n", exist, slugId)

	if !exist {
		log.Println("[VerifyOtpForSignup] signup context missing")

		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
		return
	}

	log.Printf("[VerifyOtpForSignup] parsed slug=%s\n", slugId)

	// 4. Bind request
	var req org_signup_dtos.VerifyOtpRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf(
			"[VerifyOtpForSignup] request binding failed err=%v\n",
			err,
		)

		utils.ValidationError(ctx, err)
		return
	}

	log.Printf(
		"[VerifyOtpForSignup] otp received length=%d\n",
		len(req.Otp),
	)

	// 5. Verify OTP & create org + employee
	log.Println("[VerifyOtpForSignup] calling VerifyOtp service")

	orgID, empID, err := c.signUpSvc.VerifyOtp(
		ctx,
		req.Otp,
		slugId,
	)

	if err != nil {
		log.Printf(
			"[VerifyOtpForSignup] VerifyOtp failed err=%v\n",
			err,
		)

		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			err,
		)
		return
	}

	log.Printf(
		"[VerifyOtpForSignup] VerifyOtp success orgID=%s empID=%s\n",
		orgID,
		empID,
	)

	// 6. Issue JWT
	log.Println("[VerifyOtpForSignup] generating auth token")

	token, err := c.jwtMgr.IssueWebAuthCookieToken(
		empID,
		orgID,
	)

	if err != nil {
		log.Printf(
			"[VerifyOtpForSignup] token generation failed err=%v\n",
			err,
		)

		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			err,
		)
		return
	}

	log.Println("[VerifyOtpForSignup] token generated successfully")

	// 7. Cleanup signup cookie & set auth cookie
	log.Println("[VerifyOtpForSignup] clearing signup cookie")

	c.cookies.Clear(ctx, cookieManager.Signup)

	log.Println("[VerifyOtpForSignup] setting auth cookie")

	c.cookies.Set(ctx, cookieManager.Auth, token)
	c.cookies.Set(ctx, cookieManager.CSRF, token)

	log.Println("[VerifyOtpForSignup] cookies set successfully")

	// 8. Success response
	log.Println("[VerifyOtpForSignup] request completed successfully")

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
