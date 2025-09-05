package orgcontrollers

import "github.com/gin-gonic/gin"

// ResetOrgPassword godoc
// @Summary      Reset organization password
// @Description  Reset the password for an existing organization
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email query string true "Organization email"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]string
// @Router       /org/reset-password [get]
func (c *OrganisationController) ResetOrgPassword(ctx *gin.Context) {

}
