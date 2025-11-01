package signupcontrollers

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *SignUpController) CompleteRegistration(ctx *gin.Context) {
	var req dto.Org_Complete_Registration_REQ_DTO

	// Bind and validate input
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Extract contextual details
	orgRootEmail := ctx.GetString(middlewares.EmailContextKey)
	orgSharpnrID := ctx.GetString(middlewares.OrgIDContextKey)

	// Call service method
	resData, token, err := s.service.CompleteOrgRegistration(ctx, req, orgRootEmail, orgSharpnrID)
	if err != nil {
		errorRes := dto.Response[any]{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to complete organization registration",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		constants.TOKEN_KEY_AUTHENTICATED,
		token,
		constants.TOKEN_DURATION_LONG,
		"/",
		"",
		false,
		true,
	)

	ctx.SetCookie(
		constants.TOKEN_KEY_VERIFY_OTP,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	// Success response
	successRes := dto.Response[dto.Org_Complete_Registration_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "Organization registration completed successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      resData,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, successRes)
}
