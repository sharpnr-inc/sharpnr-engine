package signupcontrollers

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"time"

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
		ctx.JSON(http.StatusOK, dto.Response[any]{
			StatusCode: int(constants.StatusBadReq),
			Message:    "Invalid Request",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		})
		return
	}

	// Step 2: Call Signup Service
	res, token, err := c.service.InitiateOrgRegistration(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response[any]{
			StatusCode: 500,
			Message:    "failed to register organization",
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

	// Step 3: Return success response
	ctx.JSON(http.StatusOK, dto.Response[dto.Org_Signup_RESPONSE_DTO]{
		StatusCode: 200,
		Message:    "OTP sent to email",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      res, // your DTO response
		Timestamp: time.Now(),
		Token:     token,
	})
}
