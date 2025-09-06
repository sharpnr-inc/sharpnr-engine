package signupcontrollers

import "github.com/gin-gonic/gin"

// CreateOrg godoc
// @Summary Create a new organization
// @Description Registers a new organization
// @Tags Organization
// @Accept json
// @Produce json
// @Param input body dto.Org_Signup_REQUEST_DTO true "Organization signup data"
// @Success 200 {object} dto.Org_Signup_RESPONSE_DTO
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (c *SignUpController) RegisterOrg(ctx *gin.Context) {
	// resDTO, _, err := c.service.InitiateOrgRegistration()

}
