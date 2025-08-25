package orgcontrollers

import "github.com/gin-gonic/gin"

// UpdateOrgRegistration godoc
// @Summary      Update organization registration details
// @Description  Update organization details after initial registration
// @Tags         Registration
// @Accept       json
// @Produce      json
// @Param        request body dto.UpdateOrgRequest true "Update org registration payload"
// @Success      200  {object} dto.UpdateOrgResponse
// @Failure      400  {object} map[string]string
// @Router       /org/update [post]
func (c *OrganisationController) UpdateOrgRegistration(ctx *gin.Context) {

}
