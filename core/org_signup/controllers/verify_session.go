package org_signup_controllers

import (
	"errors"
	"net/http"
	organization_services "sharpnr-engine/core/organization/services"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orgSignUpController) VerifySlug(ctx *gin.Context) {

	// SignupGuardMiddleware stores the slug as a uuid.UUID (it comes straight
	// off SignupClaims.UserID), so read it back through the typed helper. A
	// bare `.(string)` assertion on the context value never succeeds.
	slug, exists := webMiddleware.GetSignupUserID(ctx)
	if !exists || slug == uuid.Nil {
		utils.Failure(
			ctx,
			http.StatusForbidden,
			utils.CodeUnauthorized,
			errors.New("signup context missing or expired"),
		)
		return
	}

	// Call service
	email, err := c.signUpSvc.VerifyInitiationRegCookie(
		ctx.Request.Context(),
		slug,
	)
	if err != nil {
		// The signup token outlives the staging record it points at, so passing
		// the guard above says nothing about whether the signup still exists.
		// That case is the normal one here and gets 410 + NOT_FOUND, which the
		// client can switch on to send the user back to /signup/initiate — it
		// is not a malformed request (400) and not an auth failure (401).
		switch {
		case errors.Is(err, organization_services.ErrSignupNotFound),
			errors.Is(err, organization_services.ErrInvalidSignupData):
			utils.Failure(
				ctx,
				http.StatusGone,
				utils.CodeNotFound,
				errors.New("signup session expired, please start again"),
			)
		default:
			utils.Failure(
				ctx,
				http.StatusInternalServerError,
				utils.CodeInternalError,
				errors.New("could not load signup session, please try again"),
			)
		}
		return
	}

	utils.Success(ctx, http.StatusOK, "Signup slug verified", email)
}
