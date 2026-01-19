package org_signup_controllers

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifyOtpForSignup(ctx *gin.Context) {

	// This is Create the Employee and Org
	slug, exist := ctx.Get(webMiddleware.SignupUserIDContextKey)
	if !exist {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
	}

	// Type assert to uuid.UUID
	sluguuid, ok := slug.(uuid.UUID)
	if !ok || slug == uuid.Nil {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("invalid signup context"),
		)
		return
	}

	c.signUpSvc.VerifyOtp(ctx, "000000", sluguuid)

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
	slug, ok := userIDVal.(uuid.UUID)
	if !ok || slug == uuid.Nil {
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
