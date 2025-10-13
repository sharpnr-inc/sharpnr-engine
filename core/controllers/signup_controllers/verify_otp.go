package signupcontrollers

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/shared/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

// VerifyRegistrationOTP godoc
// @Summary      Verify OTP for organization registration
// @Description  Verifies the OTP sent to the organization's email during signup.
// @Tags         Org SignUp
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer token"
// @Param        body  body      dto.Org_Verify_Otp_REQUEST_DTO  true  "OTP verification payload"
// @Success      200   {object}  dto.Response[dto.Org_Verify_Otp_RESPONSE_DTO]
// @Failure      400   {object}  dto.Response[any]
// @Failure      401   {object}  dto.Response[any]
// @Router       /otp-verify [post]
func (c *SignUpController) VerifyRegistrationOTP(ctx *gin.Context) {
	var req dto.Org_Verify_Otp_REQUEST_DTO

	// Step 1: Bind request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, dto.Response[any]{
			StatusCode: int(constants.StatusBadReq),
			Message:    "Invalid request payload",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		})
		return
	}

	email := ctx.GetString(middlewares.EmailContextKey)

	// Step 2: Verify OTP via service
	res, ok, token, err := c.service.VerifyOrgRegistration(ctx, email, req.OTP)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "verification failed",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
			Token:     "",
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusOK, dto.Response[any]{
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid OTP or token expired",
			Error: dto.ErrorResponse{
				Message: "invalid OTP or token expired",
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
			Token:     "",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		constants.TOKEN_KEY_VERIFY_OTP, // new cookie key
		token,                          // your JWT
		constants.TOKEN_DURATION_SHORT, // expiry time
		"/",
		"",    // domain
		false, // secure false for dev
		true,  // httpOnly
	)

	// Step 3: Remove old cookie (if exists)
	ctx.SetCookie(
		constants.TOKEN_KEY_REGISTRATION,
		"",
		-1, // negative maxAge deletes the cookie
		"/",
		"",
		false, // secure false for dev (change to true in prod)
		true,  // httpOnly
	)

	// Step 5: Success response
	ctx.JSON(http.StatusOK, dto.Response[dto.Org_Verify_Otp_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "OTP verified successfully",
		Error: dto.ErrorResponse{
			IsError: false,
		},
		Data:      res,
		Timestamp: time.Now(),
		// Token:     token, // optional, mainly for frontend debugging
	})
}
