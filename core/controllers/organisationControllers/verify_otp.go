package orgcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"

	"github.com/gin-gonic/gin"
)

// VerifyOtp godoc
// @Summary      Verify organization email with OTP
// @Description  Verifies OTP sent during registration
// @Tags         Registration
// @Accept       json
// @Produce      json
// @Param        request body dto.VerifyOtpRequest true "OTP verification payload"
// @Success      200  {object} dto.VerifyOtpResponse
// @Failure      400  {object} map[string]string
// @Router       /org/verify-email-otp [post]
func (c *OrganisationController) VerifyOtp(ctx *gin.Context) {
	var req dto.OrgVerifyOTP_Req_Dto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := ctx.GetString("email")
	if email == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: email missing"})
		return
	}

	respDto, _, err := c.service.VerifyOrgRegistration(ctx.Request.Context(), email, req.Otp)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success response
	ctx.JSON(http.StatusOK, gin.H{
		"data": respDto,
	})

}
