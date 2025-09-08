package signupcontrollers

import (
	dto "sharpnr-engine/core/dtos"

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
// @Success      200   {object}  dto.Org_Verify_Otp_RESPONSE_DTO  "OTP verified successfully"
// @Failure      400   {object}  map[string]string "Invalid request payload or verification failed"
// @Failure      401   {object}  map[string]string "Invalid OTP or token expired"
// @Router       /otp-verify [post]
func (c *SignUpController) VerifyRegistrationOTP(ctx *gin.Context) {
	var req dto.Org_Verify_Otp_REQUEST_DTO

	// 1️⃣ Bind JSON body to DTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error":   "invalid request payload",
			"details": err.Error(),
		})
		return
	}

	// 2️⃣ Call service to verify OTP
	res, ok, err := c.service.VerifyOrgRegistration(ctx, req.Token, req.OTP)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   "verification failed",
			"details": err.Error(),
		})
		return
	}

	if !ok {
		ctx.JSON(401, gin.H{
			"error": "invalid OTP or token expired",
		})
		return
	}

	// @Success 200 {object} dto.Org_Verify_Otp_RESPONSE_DTO "OTP verified successfully"
	ctx.JSON(200, res)

}
