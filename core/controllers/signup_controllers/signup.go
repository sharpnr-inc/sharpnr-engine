package signupcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"

	"github.com/gin-gonic/gin"
)

// CreateOrg godoc
// @Summary Create a new organization
// @Description Registers a new organization
// @Tags  Org SignUp
// @Accept json
// @Produce json
// @Param input body dto.Org_Signup_REQUEST_DTO true "Organization signup data"
// @Success 200 {object} dto.Org_Signup_RESPONSE_DTO
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (c *SignUpController) RegisterOrg(ctx *gin.Context) {
	var req dto.Org_Signup_REQUEST_DTO

	// Step 1: Bind JSON request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"details": err.Error(),
		})
		return
	}

	// Step 2: Call Signup Service
	res, _, err := c.service.InitiateOrgRegistration(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Step 3: Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to email",
		"token":   res.Token,
		"email":   req.Email,
	})
}
