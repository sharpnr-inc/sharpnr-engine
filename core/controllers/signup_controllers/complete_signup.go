package signupcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/shared/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *SignUpController) CompleteRegistration(ctx *gin.Context) {
	var req dto.Org_Complete_Regesistration_REQ_DTO

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

	// Success response
	successRes := dto.Response[map[string]any]{
		StatusCode: http.StatusOK,
		Message:    "Organization registration completed successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data: &map[string]any{
			"org_root_email": orgRootEmail,
			"org_sharpnr_id": orgSharpnrID,
			"status":         "completed",
			"data":           resData,
		},
		Token:     token,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, successRes)
}
