package org_signup_controllers

import (
	"errors"
	"net/http"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {

	// This is Create the Employee and Org
	slugId, exist := ctx.Get(webMiddleware.SignupUserIDContextKey)
	if !exist {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
	}

	// Type assert to uuid.UUID
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

	var req org_signup_dtos.VerifyOtpRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	_, _, err = c.signUpSvc.VerifyOtp(ctx, req.Otp, slug)
	if err != nil {
		utils.Failure(ctx, http.StatusBadGateway, utils.CodeInternalError, err)
		return
	}

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
